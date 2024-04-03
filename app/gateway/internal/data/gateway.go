package data

import (
	"context"
	"github.com/jinzhu/copier"
	imPb "kratos-im/api/im"
	socialPb "kratos-im/api/social"
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
	})
	if err != nil {
		return nil, err
	}
	return &biz.GroupPutinResp{GroupId: resp.GroupId}, nil
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
