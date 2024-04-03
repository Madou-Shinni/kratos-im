package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	pb "kratos-im/api/im"
	"kratos-im/constants"
	"kratos-im/model"
	"kratos-im/pkg/tools"

	"github.com/go-kratos/kratos/v2/log"
)

var DefaultChatLogCount int64 = 100

// IMRepo is a Greater repo.
type IMRepo interface {
	Save(ctx context.Context, chatLog model.ChatLog) error
	ListBySendTime(ctx context.Context, conversationId string, startSendTime, endSendTime, count int64) ([]*model.ChatLog, error)
	FindConversationOne(ctx context.Context, id string) (*model.Conversation, error)
	FindChatLogOne(ctx context.Context, id string) (*model.ChatLog, error)
	ListByConversationIds(ctx context.Context, ids []string) ([]*model.Conversation, error)
	ConversationsByUserId(ctx context.Context, uid string) (*model.Conversations, error)
	UpdateMsg(ctx context.Context, chatLog *model.ChatLog) error
	UpdateConversations(ctx context.Context, data *model.Conversations) error
	CreateConversation(ctx context.Context, conversation model.Conversation) error
}

// IMUsecase is a IM usecase.
type IMUsecase struct {
	repo IMRepo
	log  *log.Helper
}

// NewIMUsecase new a IM usecase.
func NewIMUsecase(repo IMRepo, logger log.Logger) *IMUsecase {
	return &IMUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *IMUsecase) CreateChatLog(ctx context.Context, data *pb.ChatLog) error {
	chatLog := model.ChatLog{
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ConversationId: data.ConversationId,
		MsgFrom:        0,
		MsgType:        constants.MType(data.MsgType),
		MsgContent:     data.MsgContent,
		SendTime:       data.SendTime,
		Status:         0,
	}
	return uc.repo.Save(ctx, chatLog)
}

func (uc *IMUsecase) GetChatLog(ctx context.Context, req *pb.GetChatLogReq) ([]*model.ChatLog, error) {
	var chatLogList = make([]*model.ChatLog, 0)

	// 根据id查询聊天记录
	if req.MsgId != "" {
		one, err := uc.repo.FindChatLogOne(ctx, req.MsgId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return chatLogList, err
			}
			return nil, err
		}
		chatLogList = append(chatLogList, one)
		return chatLogList, nil
	}

	// 根据时间分段查询聊天记录
	list, err := uc.repo.ListBySendTime(ctx, req.ConversationId, req.StartSendTime, req.EndSendTime, req.Count)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		chatLogList = append(chatLogList, item)
	}

	return chatLogList, nil
}

func (uc *IMUsecase) GetConversations(ctx context.Context, req *pb.GetConversationsReq) (*pb.GetConversationsResp, error) {
	resp := &pb.GetConversationsResp{}
	// 根据用户查询用户会话列表
	data, err := uc.repo.ConversationsByUserId(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return resp, nil
		}
		return nil, err
	}

	// 根据会话列表，查询具体的会话
	ids := make([]string, 0, len(data.ConversationList))
	for _, conversation := range data.ConversationList {
		ids = append(ids, conversation.ConversationId)
	}

	list, err := uc.repo.ListByConversationIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	var res pb.GetConversationsResp
	copier.Copy(&res, &data)

	// 计算是否存在未读消息
	for _, conversation := range list {
		if _, ok := res.ConversationList[conversation.ConversationId]; !ok {
			continue
		}

		total := res.ConversationList[conversation.ConversationId].Total
		if total < int32(conversation.Total) {
			// 有新的消息
			res.ConversationList[conversation.ConversationId].Total = int32(conversation.Total)
			// 待读消息量
			res.ConversationList[conversation.ConversationId].ToRead = int32(conversation.Total) - total
			// 有新消息显示
			res.ConversationList[conversation.ConversationId].IsShow = true
		}
	}

	return &res, nil
}

func (uc *IMUsecase) PutConversations(ctx context.Context, req *pb.PutConversationsReq) (*pb.PutConversationsResp, error) {
	data, err := uc.repo.ConversationsByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if data.ConversationList == nil {
		data.ConversationList = make(map[string]*model.Conversation)
	}

	for i, conversation := range req.ConversationList {
		var oldTotal int
		if data.ConversationList[i] != nil {
			oldTotal = data.ConversationList[i].Total
		}

		data.ConversationList[i] = &model.Conversation{
			ConversationId: conversation.ConversationId,
			ChatType:       constants.ChatType(conversation.ChatType),
			IsShow:         conversation.IsShow,
			Total:          int(conversation.Read) + oldTotal,
			Seq:            conversation.Seq,
		}
	}

	err = uc.repo.UpdateConversations(ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.PutConversationsResp{}, nil
}

func (uc *IMUsecase) CreateGroupConversation(ctx context.Context, req *pb.CreateGroupConversationReq) (*pb.CreateGroupConversationResp, error) {
	res := &pb.CreateGroupConversationResp{}
	// 查询会话是否存在
	data, err := uc.repo.FindConversationOne(ctx, req.GroupId)
	if err == nil {
		return res, nil
	}

	if !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if data != nil {
		return res, err
	}

	// 添加群会话
	err = uc.repo.CreateConversation(ctx, model.Conversation{
		ConversationId: req.GroupId,
		ChatType:       constants.ChatTypeGroup,
		Total:          0,
	})
	if err != nil {
		return nil, err
	}

	_, err = uc.SetUpUserConversation(ctx, &pb.SetUpUserConversationReq{
		SendId:   req.CreateId,
		RecvId:   req.GroupId,
		ChatType: int32(constants.ChatTypeGroup),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *IMUsecase) SetUpUserConversation(ctx context.Context, req *pb.SetUpUserConversationReq) (*pb.SetUpUserConversationResp, error) {
	switch constants.ChatType(req.ChatType) {
	case constants.ChatTypeSingle: // 单聊
		// 生成会话id
		conversationId := tools.CombineId(req.SendId, req.RecvId)
		// 查询会话是否存在
		_, err := uc.repo.FindConversationOne(ctx, conversationId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				// 会话不存在，创建会话
				err = uc.repo.CreateConversation(ctx, model.Conversation{
					ConversationId: conversationId,
					ChatType:       constants.ChatTypeSingle,
					Total:          0,
				})
				if err != nil {
					return nil, err
				}
			} else {
				// 查询会话失败
				return nil, err
			}
		}

		// 建立两者会话
		err = uc.setUpUserConversation(ctx, conversationId, req.SendId, req.RecvId, constants.ChatTypeSingle, true)
		if err != nil {
			return nil, err
		}
		err = uc.setUpUserConversation(ctx, conversationId, req.RecvId, req.SendId, constants.ChatTypeSingle, false) // 接收方不显示会话
		if err != nil {
			return nil, err
		}
	case constants.ChatTypeGroup: // 群聊
		err := uc.setUpUserConversation(ctx, req.RecvId, req.SendId, req.SendId, constants.ChatTypeGroup, false) // 接收方不显示会话
		if err != nil {
			return nil, err
		}
	}

	return &pb.SetUpUserConversationResp{}, nil
}

// SetUpUserConversation 设置用户会话
func (uc *IMUsecase) setUpUserConversation(ctx context.Context, conversationId, userId, recvId string, chatType constants.ChatType, isShow bool) error {
	// 查询用户会话列表
	data, err := uc.repo.ConversationsByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// 用户会话列表不存在
			data = &model.Conversations{
				ID:               primitive.NewObjectID(),
				UserId:           userId,
				ConversationList: make(map[string]*model.Conversation),
			}
		} else {
			return err
		}
	}

	// 创建用户会话列表
	if _, ok := data.ConversationList[conversationId]; ok {
		// 会话已存在
		return nil
	}

	// 会话不存在，创建会话
	data.ConversationList[conversationId] = &model.Conversation{
		ConversationId: conversationId,
		ChatType:       chatType,
		IsShow:         isShow,
	}

	// 更新用户会话列表
	err = uc.repo.UpdateConversations(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
