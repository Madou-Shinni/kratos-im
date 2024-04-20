package biz

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-im/api/gateway"
	imPb "kratos-im/api/im"
	socialPb "kratos-im/api/social"
	userPb "kratos-im/api/user"
	"kratos-im/common"
	"kratos-im/constants"
	"kratos-im/model"
	"kratos-im/pkg/rws"
	"kratos-im/pkg/tools"
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

type GroupPutInHandleReq struct {
	GroupReqId   uint64
	HandleUid    string
	HandleResult constants.HandleResult
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
	GroupPutInHandle(ctx context.Context, data *GroupPutInHandleReq) error
	GroupPutinList(ctx context.Context, groupId uint64) ([]*model.GroupRequests, error)
	GroupList(ctx context.Context, userId string) ([]*model.Groups, error)
	GroupUsers(ctx context.Context, groupId uint64) ([]*model.GroupMembers, error)
	GetChatLog(ctx context.Context, req *imPb.GetChatLogReq) ([]*model.ChatLog, error)
	UserLogin(ctx context.Context, data *userPb.LoginRequest) (*userPb.LoginReply, error)
	HSetOnlineUser(ctx context.Context, userId string, status bool) error
	GetOnlineUser(ctx context.Context) (map[string]string, error)
	GetConversations(ctx context.Context, req *imPb.GetConversationsReq) (*imPb.GetConversationsResp, error)
	ListUserByIds(ctx context.Context, ids []string) (*userPb.ListResp, error)
	ListGroupByIds(ctx context.Context, ids []uint64) (*socialPb.GroupMapResp, error)
	PutConversations(ctx context.Context, req *imPb.PutConversationsReq) (*imPb.PutConversationsResp, error)
	UserSignUp(ctx context.Context, data *userPb.Account) error
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

// GroupPutin 入群申请
func (uc *GatewayUsecase) GroupPutin(ctx context.Context, req *v1.GroupPutinReq) (*v1.GroupPutinResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 入群申请
	data, err := uc.repo.GroupPutin(ctx, &model.GroupRequests{
		GroupId:    req.GroupId,
		ReqId:      uid,
		ReqMsg:     req.ReqMsg,
		JoinSource: int(req.JoinSource),
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
			FriendUid: v.FriendUid,
			Remark:    v.Remark,
		})
	}

	return &v1.FriendListResp{List: list}, nil
}

// GroupPutInHandle 入群申请处理
func (uc *GatewayUsecase) GroupPutInHandle(ctx context.Context, req *v1.GroupPutInHandleReq) (*v1.GroupPutInHandleResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = uc.repo.GroupPutInHandle(ctx, &GroupPutInHandleReq{
		GroupReqId:   uint64(req.GroupReqId),
		HandleUid:    uid,
		HandleResult: constants.HandleResult(req.HandleResult),
	})
	if err != nil {
		return nil, err
	}

	if req.HandleResult == int32(constants.HandleResultAgree) {
		// 建立会话
		err = uc.repo.CreateConversation(ctx, &CreateConversationReq{
			UserId:   uid,
			RecvId:   strconv.FormatUint(req.GroupId, 10),
			ChatType: constants.ChatTypeGroup,
		})
		if err != nil {
			return nil, err
		}
	}

	return &v1.GroupPutInHandleResp{}, nil
}

// GroupPutinList 入群申请列表
func (uc *GatewayUsecase) GroupPutinList(ctx context.Context, req *v1.GroupPutinListReq) (*v1.GroupPutinListResp, error) {
	data, err := uc.repo.GroupPutinList(ctx, req.GroupId)

	if err != nil {
		return nil, err
	}

	var list = make([]*v1.GroupRequests, 0, len(data))
	err = copier.Copy(&list, &data)
	if err != nil {
		return nil, err
	}

	return &v1.GroupPutinListResp{List: list}, nil
}

// GroupList 群列表
func (uc *GatewayUsecase) GroupList(ctx context.Context, req *v1.GroupListReq) (*v1.GroupListResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	data, err := uc.repo.GroupList(ctx, uid)
	if err != nil {
		return nil, err
	}

	var list = make([]*v1.Groups, 0, len(data))
	err = copier.Copy(&list, &data)
	if err != nil {
		return nil, err
	}

	return &v1.GroupListResp{List: list}, nil
}

// GroupUserList 群成员列表
func (uc *GatewayUsecase) GroupUserList(ctx context.Context, req *v1.GroupUsersReq) (*v1.GroupUsersResp, error) {
	data, err := uc.repo.GroupUsers(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	var list = make([]*v1.GroupMembers, 0, len(data))
	err = copier.Copy(&list, &data)
	if err != nil {
		return nil, err
	}

	return &v1.GroupUsersResp{List: list}, nil
}

func (uc *GatewayUsecase) GetReadChatRecords(ctx context.Context, req *v1.GetReadChatRecordsReq) (*v1.GetReadChatRecordsResp, error) {
	// 查询消息记录
	chatLogs, err := uc.repo.GetChatLog(ctx, &imPb.GetChatLogReq{
		MsgId: req.MsgId,
	})
	if err != nil {
		return nil, err
	}

	if len(chatLogs) == 0 {
		return &v1.GetReadChatRecordsResp{}, nil
	}

	var (
		chatLog = chatLogs[0]
		reads   = []string{chatLog.SendId}
		unReads []string
	)

	// 设置未读已读
	switch chatLog.ChatType {
	case constants.ChatTypeSingle: // 单聊
		if len(chatLog.ReadRecords) == 0 || chatLog.ReadRecords[0] == 0 {
			unReads = []string{chatLog.RecvId}
		} else {
			reads = append(reads, chatLog.RecvId)
		}
	case constants.ChatTypeGroup: // 群聊
		gid, _ := strconv.ParseUint(chatLog.RecvId, 10, 64)
		members, err := uc.repo.GroupUsers(ctx, gid)
		if err != nil {
			return nil, err
		}

		bitmaps := tools.Load(chatLog.ReadRecords)
		for _, v := range members {
			if v.UserId == chatLog.SendId {
				continue
			}

			if bitmaps.IsSet(v.UserId) {
				reads = append(reads, v.UserId)
			} else {
				unReads = append(unReads, v.UserId)
			}
		}
	}

	return &v1.GetReadChatRecordsResp{
		Reads:   reads,
		UnReads: unReads,
	}, err
}

func (uc *GatewayUsecase) UserLogin(ctx context.Context, req *userPb.LoginRequest) (*v1.UserLoginResp, error) {
	uc.log.Infof("req: type: %v, payload: %v", req.Type, req.Payload)
	// 用户登录
	resp, err := uc.repo.UserLogin(ctx, req)
	if err != nil {
		return nil, err
	}

	// 缓存登录状态
	err = uc.repo.HSetOnlineUser(ctx, resp.UserInfo.UserId, true)
	if err != nil {
		return nil, err
	}

	return &v1.UserLoginResp{
		UserInfo: &v1.UserLoginResp_UserInfo{
			UserId:   resp.UserInfo.UserId,
			Avatar:   resp.UserInfo.Avatar,
			Nickname: resp.UserInfo.Nickname,
		},
		Token: resp.Token,
	}, nil
}

func (uc *GatewayUsecase) FriendsOnline(ctx context.Context, req *v1.FriendsOnlineReq) (*v1.FriendsOnlineResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 查询好友列表
	list, err := uc.repo.FriendList(ctx, uid)
	if err != nil {
		return nil, err
	}

	if list == nil || len(list) == 0 {
		return &v1.FriendsOnlineResp{}, nil
	}

	// 查询在线用户
	usersOnline, err := uc.repo.GetOnlineUser(ctx)
	if err != nil {
		return nil, err
	}

	// 过滤在线好友
	friendsOnline := make(map[string]bool, len(list))
	for _, v := range list {
		if _, ok := usersOnline[v.FriendUid]; ok {
			friendsOnline[v.FriendUid] = true
		} else {
			friendsOnline[v.FriendUid] = false
		}
	}

	return &v1.FriendsOnlineResp{OnlineList: friendsOnline}, nil
}

func (uc *GatewayUsecase) GroupMembersOnline(ctx context.Context, req *v1.GroupMembersOnlineReq) (*v1.GroupMembersOnlineResp, error) {
	// 查询群成员列表
	list, err := uc.repo.GroupUsers(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	if list == nil || len(list) == 0 {
		return &v1.GroupMembersOnlineResp{}, nil
	}

	// 查询在线用户
	usersOnline, err := uc.repo.GetOnlineUser(ctx)
	if err != nil {
		return nil, err
	}

	// 过滤在线群成员
	membersOnline := make(map[string]bool, len(list))
	for _, v := range list {
		if _, ok := usersOnline[v.UserId]; ok {
			membersOnline[v.UserId] = true
		} else {
			membersOnline[v.UserId] = false
		}
	}

	return &v1.GroupMembersOnlineResp{OnlineList: membersOnline}, nil
}

func (uc *GatewayUsecase) GetConversations(ctx context.Context, req *v1.GetConversationsReq) (*v1.GetConversationsResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	conversations, err := uc.repo.GetConversations(ctx, &imPb.GetConversationsReq{UserId: uid})
	if err != nil {
		return nil, err
	}

	var (
		groupIds = make([]uint64, 0, len(conversations.ConversationList))
		userIds  = make([]string, 0, len(conversations.ConversationList))
	)

	// 查询用户和群信息，填充会话内容
	for _, v := range conversations.ConversationList {
		if v.ChatType == constants.ChatTypeSingle {
			userIds = append(userIds, v.TargetId)
		} else {
			gid, _ := strconv.ParseUint(v.ConversationId, 10, 64)
			groupIds = append(groupIds, gid)
		}
	}

	// 查询用户信息
	usersResp, err := uc.repo.ListUserByIds(ctx, userIds)
	if err != nil {
		return nil, err
	}

	// 查询群信息
	groupsResp, err := uc.repo.ListGroupByIds(ctx, groupIds)
	if err != nil {
		return nil, err
	}

	var res v1.GetConversationsResp
	res.Conversations = make(map[string]*v1.Conversation, len(conversations.ConversationList))

	for k, v := range conversations.ConversationList {
		var nickname, avatar string // 填充会话内容

		switch v.ChatType {
		case constants.ChatTypeSingle:
			if user, ok := usersResp.Users[v.TargetId]; ok {
				nickname = user.Nickname
				avatar = user.Avatar
			}
		case constants.ChatTypeGroup:
			tid, _ := strconv.ParseUint(v.ConversationId, 10, 64)
			if group, ok := groupsResp.GroupMap[tid]; ok {
				nickname = group.Name
				avatar = group.Icon
			}
		}

		res.Conversations[k] = &v1.Conversation{
			ConversationId: v.ConversationId,
			ChatType:       v.ChatType,
			Total:          v.Total,
			Read:           v.Total - v.ToRead,
			Unread:         v.ToRead,
			IsShow:         v.IsShow,
			Seq:            v.Seq,
			TargetId:       v.TargetId,
			Nickname:       nickname,
			Avatar:         avatar,
		}
	}

	res.UserId = uid

	return &res, nil
}

func (uc *GatewayUsecase) GetChatLog(ctx context.Context, req *v1.GetChatLogReq) (*v1.GetChatLogResp, error) {
	chatLogs, err := uc.repo.GetChatLog(ctx, &imPb.GetChatLogReq{
		ConversationId: req.ConversationId,
		StartSendTime:  req.StartSendTime,
		EndSendTime:    req.EndSendTime,
		Count:          req.Count,
	})
	if err != nil {
		return nil, err
	}

	var list = make([]*v1.ChatLog, 0, len(chatLogs))
	err = copier.Copy(&list, &chatLogs)
	if err != nil {
		return nil, err
	}

	return &v1.GetChatLogResp{List: list}, nil
}

func (uc *GatewayUsecase) PutConversations(ctx context.Context, req *v1.PutConversationsReq) (*v1.PutConversationsResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 更新会话
	var conversationList = make(map[string]*imPb.Conversation, len(req.Conversations))
	err = copier.Copy(&conversationList, &req.Conversations)
	if err != nil {
		return nil, err
	}

	_, err = uc.repo.PutConversations(ctx, &imPb.PutConversationsReq{
		UserId:           uid,
		ConversationList: conversationList,
	})
	if err != nil {
		return nil, err
	}

	return &v1.PutConversationsResp{}, nil
}

func (uc *GatewayUsecase) SetUpUserConversation(ctx context.Context, req *v1.SetUpUserConversationReq) (*v1.SetUpUserConversationResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 建立会话
	err = uc.repo.CreateConversation(ctx, &CreateConversationReq{
		UserId:   uid,
		RecvId:   req.RecvId,
		ChatType: constants.ChatType(req.ChatType),
	})
	if err != nil {
		return nil, err
	}

	return &v1.SetUpUserConversationResp{}, nil
}

func (uc *GatewayUsecase) UserSignUp(ctx context.Context, req *userPb.Account) (*emptypb.Empty, error) {
	// 用户注册
	err := uc.repo.UserSignUp(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
