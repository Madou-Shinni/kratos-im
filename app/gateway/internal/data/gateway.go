package data

import (
	"context"
	"github.com/jinzhu/copier"
	imPb "kratos-im/api/im"
	socialPb "kratos-im/api/social"
	userPb "kratos-im/api/user"
	"kratos-im/constants"
	"kratos-im/model"
	"strconv"

	"kratos-im/app/gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayRepo struct {
	data *Data
	log  *log.Helper
}

// NewGatewayRepo .
func NewGatewayRepo(data *Data, logger log.Logger) biz.GatewayRepo {
	return &gatewayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gatewayRepo) Save(ctx context.Context, data model.ChatLog) error {
	_, err := r.data.imClient.CreateChatLog(ctx, &imPb.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		MsgType:        int32(data.MsgType),
		MsgContent:     data.MsgContent,
		ChatType:       int32(data.ChatType),
		SendTime:       data.SendTime,
	})
	if err != nil {
		return err
	}
	return nil
}

// GroupCreate 创建群
func (r *gatewayRepo) GroupCreate(ctx context.Context, data *model.Groups) (*biz.GroupCreateResp, error) {
	resp, err := r.data.socialClient.GroupCreate(ctx, &socialPb.GroupCreateReq{
		Name:       data.Name,
		Icon:       data.Icon,
		Status:     int32(data.Status),
		CreatorUid: data.CreatorUid,
	})
	if err != nil {
		return nil, err
	}
	return &biz.GroupCreateResp{Id: resp.Id}, nil
}

// GroupPutin 入群申请
func (r *gatewayRepo) GroupPutin(ctx context.Context, data *model.GroupRequests) (*biz.GroupPutinResp, error) {
	resp, err := r.data.socialClient.GroupPutin(ctx, &socialPb.GroupPutinReq{
		GroupId:    data.GroupId,
		ReqId:      data.ReqId,
		ReqMsg:     data.ReqMsg,
		InviterUid: data.InviterUserId,
		JoinSource: int32(data.JoinSource),
	})
	if err != nil {
		return nil, err
	}
	return &biz.GroupPutinResp{GroupId: resp.GroupId}, nil
}

// GroupPutInHandle 入群申请处理
func (r *gatewayRepo) GroupPutInHandle(ctx context.Context, data *biz.GroupPutInHandleReq) error {
	_, err := r.data.socialClient.GroupPutInHandle(ctx, &socialPb.GroupPutInHandleReq{
		GroupReqId:   int32(data.GroupReqId),
		HandleUid:    data.HandleUid,
		HandleResult: int32(data.HandleResult),
	})
	if err != nil {
		return err
	}
	return nil
}

// GroupPutinList 入群申请列表
func (r *gatewayRepo) GroupPutinList(ctx context.Context, groupId uint64) ([]*model.GroupRequests, error) {
	resp, err := r.data.socialClient.GroupPutinList(ctx, &socialPb.GroupPutinListReq{
		GroupId: groupId,
	})
	if err != nil {
		return nil, err
	}
	var data []*model.GroupRequests

	err = copier.Copy(&data, &resp.List)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GroupList 群列表
func (r *gatewayRepo) GroupList(ctx context.Context, userId string) ([]*model.Groups, error) {
	resp, err := r.data.socialClient.GroupList(ctx, &socialPb.GroupListReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	var data []*model.Groups

	err = copier.Copy(&data, &resp.List)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GroupUsers 群成员列表
func (r *gatewayRepo) GroupUsers(ctx context.Context, groupId uint64) ([]*model.GroupMembers, error) {
	resp, err := r.data.socialClient.GroupUsers(ctx, &socialPb.GroupUsersReq{
		GroupId: groupId,
	})
	if err != nil {
		return nil, err
	}
	var data []*model.GroupMembers

	err = copier.Copy(&data, &resp.List)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateGroupConversation 创建群会话
func (r *gatewayRepo) CreateGroupConversation(ctx context.Context, data *biz.GroupConversationReq) error {
	_, err := r.data.imClient.CreateGroupConversation(ctx, &imPb.CreateGroupConversationReq{
		GroupId:  strconv.FormatUint(data.GroupId, 10),
		CreateId: data.CreatedId,
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateConversation 创建会话(按申请情况自动处理)
func (r *gatewayRepo) CreateConversation(ctx context.Context, data *biz.CreateConversationReq) error {
	_, err := r.data.imClient.SetUpUserConversation(ctx, &imPb.SetUpUserConversationReq{
		SendId:   data.UserId,
		RecvId:   data.RecvId,
		ChatType: int32(data.ChatType),
	})
	if err != nil {
		return err
	}
	return nil
}

// FriendPutIn 好友申请
func (r *gatewayRepo) FriendPutIn(ctx context.Context, data *model.FriendRequests) error {
	_, err := r.data.socialClient.FriendPutIn(ctx, &socialPb.FriendPutInReq{
		UserId: data.UserId,
		ReqUid: data.ReqUid,
		ReqMsg: data.ReqMsg,
	})
	if err != nil {
		return err
	}
	return nil
}

// FriendPutInHandle 好友申请处理
func (r *gatewayRepo) FriendPutInHandle(ctx context.Context, data *biz.FriendPutInHandleReq) error {
	_, err := r.data.socialClient.FriendPutInHandle(ctx, &socialPb.FriendPutInHandleReq{
		FriendReqId:  data.FriendReqId,
		UserId:       data.UserId,
		HandleResult: data.HandleResult,
	})
	if err != nil {
		return err
	}
	return nil
}

// FriendPutInList 好友申请列表
func (r *gatewayRepo) FriendPutInList(ctx context.Context, userId string) ([]*model.FriendRequests, error) {
	resp, err := r.data.socialClient.FriendPutInList(ctx, &socialPb.FriendPutInListReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	var data []*model.FriendRequests

	err = copier.Copy(&data, &resp.List)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FriendList 好友列表
func (r *gatewayRepo) FriendList(ctx context.Context, userId string) ([]*model.Friends, error) {
	resp, err := r.data.socialClient.FriendList(ctx, &socialPb.FriendListReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	var data []*model.Friends

	err = copier.Copy(&data, &resp.List)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetChatLog 获取聊天记录
func (r *gatewayRepo) GetChatLog(ctx context.Context, req *imPb.GetChatLogReq) ([]*model.ChatLog, error) {
	resp, err := r.data.imClient.GetChatLog(ctx, req)
	if err != nil {
		return nil, err
	}
	var data []*model.ChatLog

	err = copier.Copy(&data, &resp.List)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UserLogin 用户登录
func (r *gatewayRepo) UserLogin(ctx context.Context, data *userPb.LoginRequest) (*userPb.LoginReply, error) {
	resp, err := r.data.userClient.Login(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HSetOnlineUser 缓存在线状态
func (r *gatewayRepo) HSetOnlineUser(ctx context.Context, userId string, status bool) error {
	return r.data.rdb.HSet(ctx, constants.OnlineUserKey, userId, status).Err()
}

// GetOnlineUser 查询在线用户
func (r *gatewayRepo) GetOnlineUser(ctx context.Context) (map[string]string, error) {
	return r.data.rdb.HGetAll(ctx, constants.OnlineUserKey).Result()
}

// GetConversations 获取会话列表
func (r *gatewayRepo) GetConversations(ctx context.Context, req *imPb.GetConversationsReq) (*imPb.GetConversationsResp, error) {
	resp, err := r.data.imClient.GetConversations(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PutConversations 更新会话
func (r *gatewayRepo) PutConversations(ctx context.Context, req *imPb.PutConversationsReq) (*imPb.PutConversationsResp, error) {
	resp, err := r.data.imClient.PutConversations(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListUserByIds 查询用户列表
func (r *gatewayRepo) ListUserByIds(ctx context.Context, ids []string) (*userPb.ListResp, error) {
	resp, err := r.data.userClient.List(ctx, &userPb.ListRequest{Ids: ids})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListGroupByIds 查询群列表
func (r *gatewayRepo) ListGroupByIds(ctx context.Context, ids []uint64) (*socialPb.GroupMapResp, error) {
	resp, err := r.data.socialClient.GroupMap(ctx, &socialPb.GroupMapReq{GroupIds: ids})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
