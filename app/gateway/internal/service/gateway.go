package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	pb "kratos-im/api/gateway"
	v1 "kratos-im/api/gateway"
	"kratos-im/app/gateway/internal/biz"
)

// GatewayService is a greeter service.
type GatewayService struct {
	v1.UnimplementedGatewayServer
	uc          *biz.GatewayUsecase
	KafkaBroker broker.Broker
}

// NewGatewayService new a greeter service.
func NewGatewayService(uc *biz.GatewayUsecase) *GatewayService {
	return &GatewayService{uc: uc}
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
