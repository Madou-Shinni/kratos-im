package service

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tx7do/kratos-transport/broker"
	pb "kratos-im/api/gateway"
	v1 "kratos-im/api/gateway"
	"kratos-im/app/gateway/internal/biz"
	"kratos-im/pkg/tools"
	"math"
	"time"
)

// GatewayService is a greeter service.
type GatewayService struct {
	v1.UnimplementedGatewayServer
	uc          *biz.GatewayUsecase
	KafkaBroker broker.Broker
}

// NewGatewayService new a greeter service.
func NewGatewayService(uc *biz.GatewayUsecase, KafkaBroker broker.Broker) *GatewayService {
	return &GatewayService{uc: uc, KafkaBroker: KafkaBroker}
}

// GroupPutin 入群申请
func (s *GatewayService) GroupPutin(ctx context.Context, req *pb.GroupPutinReq) (*pb.GroupPutinResp, error) {
	return s.uc.GroupPutin(ctx, req)
}

// GroupCreate 创建群
func (s *GatewayService) GroupCreate(ctx context.Context, req *pb.GroupCreateReq) (*pb.GroupCreateResp, error) {
	return s.uc.GroupCreate(ctx, req)
}

// GroupPutInHandle 入群申请处理
func (s *GatewayService) GroupPutInHandle(ctx context.Context, req *pb.GroupPutInHandleReq) (*pb.GroupPutInHandleResp, error) {
	return s.uc.GroupPutInHandle(ctx, req)
}

// GroupPutinList 入群申请列表
func (s *GatewayService) GroupPutinList(ctx context.Context, req *pb.GroupPutinListReq) (*pb.GroupPutinListResp, error) {
	return s.uc.GroupPutinList(ctx, req)
}

// GroupList 群列表
func (s *GatewayService) GroupList(ctx context.Context, req *pb.GroupListReq) (*pb.GroupListResp, error) {
	return s.uc.GroupList(ctx, req)
}

// GroupUsers 群成员列表
func (s *GatewayService) GroupUsers(ctx context.Context, req *pb.GroupUsersReq) (*pb.GroupUsersResp, error) {
	return s.uc.GroupUserList(ctx, req)
}

// FriendPutIn 好友申请
func (s *GatewayService) FriendPutIn(ctx context.Context, req *pb.FriendPutInReq) (*pb.FriendPutInResp, error) {
	return s.uc.FriendPutIn(ctx, req)
}

// FriendPutInHandle 好友申请处理
func (s *GatewayService) FriendPutInHandle(ctx context.Context, req *pb.FriendPutInHandleReq) (*pb.FriendPutInHandleResp, error) {
	return s.uc.FriendPutInHandle(ctx, req)
}

// FriendPutInList 好友申请列表
func (s *GatewayService) FriendPutInList(ctx context.Context, req *pb.FriendPutInListReq) (*pb.FriendPutInListResp, error) {
	return s.uc.FriendPutInList(ctx, req)
}

// FriendList 好友列表
func (s *GatewayService) FriendList(ctx context.Context, req *pb.FriendListReq) (*pb.FriendListResp, error) {
	return s.uc.FriendList(ctx, req)
}

// GetReadChatRecords 获取已读消息记录
func (s *GatewayService) GetReadChatRecords(ctx context.Context, req *pb.GetReadChatRecordsReq) (*pb.GetReadChatRecordsResp, error) {
	return s.uc.GetReadChatRecords(ctx, req)
}

// UserLogin 用户登录
func (s *GatewayService) UserLogin(ctx context.Context, req *pb.UserLoginReq) (*pb.UserLoginResp, error) {
	return s.uc.UserLogin(ctx, req)
}

// FriendsOnline 在线好友情况
func (s *GatewayService) FriendsOnline(ctx context.Context, req *pb.FriendsOnlineReq) (*pb.FriendsOnlineResp, error) {
	return s.uc.FriendsOnline(ctx, req)
}

// GroupMembersOnline 在线群成员情况
func (s *GatewayService) GroupMembersOnline(ctx context.Context, req *pb.GroupMembersOnlineReq) (*pb.GroupMembersOnlineResp, error) {
	return s.uc.GroupMembersOnline(ctx, req)
}

// GetConversations 获取会话列表
func (s *GatewayService) GetConversations(ctx context.Context, req *pb.GetConversationsReq) (*pb.GetConversationsResp, error) {
	return s.uc.GetConversations(ctx, req)
}

// GetChatLog 获取聊天记录
func (s *GatewayService) GetChatLog(ctx context.Context, req *pb.GetChatLogReq) (*pb.GetChatLogResp, error) {
	return s.uc.GetChatLog(ctx, req)
}

// PutConversations 更新会话
func (s *GatewayService) PutConversations(ctx context.Context, req *pb.PutConversationsReq) (*pb.PutConversationsResp, error) {
	return s.uc.PutConversations(ctx, req)
}

// SetUpUserConversation 建立会话
func (s *GatewayService) SetUpUserConversation(ctx context.Context, req *pb.SetUpUserConversationReq) (*pb.SetUpUserConversationResp, error) {
	return s.uc.SetUpUserConversation(ctx, req)
}

// GenToken 生成token
func (s *GatewayService) GenToken(secret string) (string, error) {
	return tools.GenToken(jwt.MapClaims{
		"userId": fmt.Sprint("kratos-im:server-discover:", time.Now().UnixMilli()),
	}, math.MaxInt,
		secret)
}
