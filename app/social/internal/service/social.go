package service

import (
	"context"

	pb "kratos-im/api/social"
)

type SocialService struct {
	pb.UnimplementedSocialServer
}

func NewSocialService() *SocialService {
	return &SocialService{}
}

func (s *SocialService) FriendPutIn(ctx context.Context, req *pb.FriendPutInReq) (*pb.FriendPutInResp, error) {
	return &pb.FriendPutInResp{}, nil
}
func (s *SocialService) FriendPutInHandle(ctx context.Context, req *pb.FriendPutInHandleReq) (*pb.FriendPutInHandleResp, error) {
	return &pb.FriendPutInHandleResp{}, nil
}
func (s *SocialService) FriendPutInList(ctx context.Context, req *pb.FriendPutInListReq) (*pb.FriendPutInListResp, error) {
	return &pb.FriendPutInListResp{}, nil
}
func (s *SocialService) FriendList(ctx context.Context, req *pb.FriendListReq) (*pb.FriendListResp, error) {
	return &pb.FriendListResp{}, nil
}
func (s *SocialService) GroupCreate(ctx context.Context, req *pb.GroupCreateReq) (*pb.GroupCreateResp, error) {
	return &pb.GroupCreateResp{}, nil
}
func (s *SocialService) GroupPutin(ctx context.Context, req *pb.GroupPutinReq) (*pb.GroupPutinResp, error) {
	return &pb.GroupPutinResp{}, nil
}
func (s *SocialService) GroupPutinList(ctx context.Context, req *pb.GroupPutinListReq) (*pb.GroupPutinListResp, error) {
	return &pb.GroupPutinListResp{}, nil
}
func (s *SocialService) GroupPutInHandle(ctx context.Context, req *pb.GroupPutInHandleReq) (*pb.GroupPutInHandleResp, error) {
	return &pb.GroupPutInHandleResp{}, nil
}
func (s *SocialService) GroupList(ctx context.Context, req *pb.GroupListReq) (*pb.GroupListResp, error) {
	return &pb.GroupListResp{}, nil
}
func (s *SocialService) GroupUsers(ctx context.Context, req *pb.GroupUsersReq) (*pb.GroupUsersResp, error) {
	return &pb.GroupUsersResp{}, nil
}
