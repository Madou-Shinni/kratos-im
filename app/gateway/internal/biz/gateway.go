package biz

import (
	"context"
	v1 "kratos-im/api/gateway"
	"kratos-im/common"
	"kratos-im/constants"
	"kratos-im/model"
	"kratos-im/pkg/rws"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

// GroupCreateResp is a response.
type GroupCreateResp struct {
	Id uint64 // 群id
}

// GroupPutinResp is a response.
type GroupPutinResp struct {
	GroupId uint64 // 群id
}

// GroupConversationReq is a request.
type GroupConversationReq struct {
	GroupId   uint64 // 群id
	CreatedId string // 创建者id
}

// CreateConversationReq is a request.
type CreateConversationReq struct {
	UserId   string             // 用户id
	RecvId   string             // 接收者id
	ChatType constants.ChatType // 聊天类型
}

type FriendPutInHandleReq struct {
	FriendReqId  int32
	UserId       string
	HandleResult int32
}

// GatewayRepo is a Greater repo.
type GatewayRepo interface {
	Save(ctx context.Context, chatLog model.ChatLog) error
	GroupCreate(ctx context.Context, data *model.Groups) (*GroupCreateResp, error)
	GroupPutin(ctx context.Context, data *model.GroupRequests) (*GroupPutinResp, error)
	CreateGroupConversation(ctx context.Context, data *GroupConversationReq) error
	CreateConversation(ctx context.Context, data *CreateConversationReq) error
	FriendPutIn(ctx context.Context, data *model.FriendRequests) error
	FriendPutInHandle(ctx context.Context, data *FriendPutInHandleReq) error
	FriendPutInList(ctx context.Context, userId string) ([]*model.FriendRequests, error)
	FriendList(ctx context.Context, userId string) ([]*model.Friends, error)
}

// GatewayUsecase is a Gateway usecase.
type GatewayUsecase struct {
	repo GatewayRepo
	log  *log.Helper
}

// NewGatewayUsecase new a Gateway usecase.
func NewGatewayUsecase(repo GatewayRepo, logger log.Logger) *GatewayUsecase {
	return &GatewayUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GatewayUsecase) CreateChatLog(ctx context.Context, data *rws.Chat) error {
	chatLog := model.ChatLog{
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ConversationId: data.ConversationId,
		MsgFrom:        0,
		MsgType:        data.Msg.MType,
		MsgContent:     data.Msg.Content,
		SendTime:       data.SendTime,
		Status:         0,
	}
	return uc.repo.Save(ctx, chatLog)
}

func (uc *GatewayUsecase) GroupPutin(ctx context.Context, req *v1.GroupPutinReq) (*v1.GroupPutinResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 入群申请
	data, err := uc.repo.GroupPutin(ctx, &model.GroupRequests{
		GroupId:       req.GroupId,
		ReqId:         uid,
		ReqMsg:        req.ReqMsg,
		InviterUserId: req.InviterUid,
	})
	if err != nil {
		return nil, err
	}

	if data.GroupId == 0 {
		return &v1.GroupPutinResp{}, nil
	}

	// 建立会话
	err = uc.repo.CreateConversation(ctx, &CreateConversationReq{
		UserId:   uid,
		RecvId:   strconv.FormatUint(data.GroupId, 10),
		ChatType: constants.ChatTypeGroup,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GroupPutinResp{GroupId: data.GroupId}, nil
}

// GroupCreate 创建群
func (uc *GatewayUsecase) GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (*v1.GroupCreateResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 创建群
	data, err := uc.repo.GroupCreate(ctx, &model.Groups{
		Name:       req.Name,
		Icon:       req.Icon,
		Status:     int(req.Status),
		CreatorUid: uid,
	})
	if err != nil {
		return nil, err
	}

	// 建立会话
	err = uc.repo.CreateGroupConversation(ctx, &GroupConversationReq{
		GroupId:   data.Id,
		CreatedId: uid,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GroupCreateResp{Id: data.Id}, nil
}

// FriendPutIn 好友申请
func (uc *GatewayUsecase) FriendPutIn(ctx context.Context, req *v1.FriendPutInReq) (*v1.FriendPutInResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 好友申请
	err = uc.repo.FriendPutIn(ctx, &model.FriendRequests{
		UserId: uid,
		ReqUid: req.ReqUid,
		ReqMsg: req.ReqMsg,
	})
	if err != nil {
		return nil, err
	}

	return &v1.FriendPutInResp{}, nil
}

// FriendPutInHandle 好友申请处理
func (uc *GatewayUsecase) FriendPutInHandle(ctx context.Context, req *v1.FriendPutInHandleReq) (*v1.FriendPutInHandleResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 好友申请处理
	err = uc.repo.FriendPutInHandle(ctx, &FriendPutInHandleReq{
		FriendReqId:  req.FriendReqId,
		UserId:       uid,
		HandleResult: req.HandleResult,
	})
	if err != nil {
		return nil, err
	}

	return &v1.FriendPutInHandleResp{}, nil
}

// FriendPutInList 好友申请列表
func (uc *GatewayUsecase) FriendPutInList(ctx context.Context, req *v1.FriendPutInListReq) (*v1.FriendPutInListResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 好友申请列表
	data, err := uc.repo.FriendPutInList(ctx, uid)
	if err != nil {
		return nil, err
	}

	var list = make([]*v1.FriendRequests, 0, len(data))
	for _, v := range data {
		list = append(list, &v1.FriendRequests{
			Id:           int32(v.ID),
			UserId:       v.UserId,
			ReqUid:       v.ReqUid,
			ReqMsg:       v.ReqMsg,
			ReqTime:      v.ReqTime.Unix(),
			HandleResult: int32(v.HandleResult),
		})
	}

	return &v1.FriendPutInListResp{List: list}, nil
}

// FriendList 好友列表
func (uc *GatewayUsecase) FriendList(ctx context.Context, req *v1.FriendListReq) (*v1.FriendListResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 好友列表
	data, err := uc.repo.FriendList(ctx, uid)
	if err != nil {
		return nil, err
	}

	var list = make([]*v1.Friends, 0, len(data))
	for _, v := range data {
		list = append(list, &v1.Friends{
			Id:        int32(v.ID),
			UserId:    v.UserId,
			FriendUid: v.FriendUid,
			Remark:    v.Remark,
		})
	}

	return &v1.FriendListResp{List: list}, nil
}
