package service

import (
	"context"
	"github.com/jinzhu/copier"
	"kratos-im/app/social/internal/biz"

	pb "kratos-im/api/social"
)

type SocialService struct {
	pb.UnimplementedSocialServer
	uc *biz.SocialUsecase
}

func NewSocialService(uc *biz.SocialUsecase) *SocialService {
	return &SocialService{uc: uc}
}

func (s *SocialService) FriendPutIn(ctx context.Context, req *pb.FriendPutInReq) (*pb.FriendPutInResp, error) {
	return s.uc.FriendPutIn(ctx, req)
}
func (s *SocialService) FriendPutInHandle(ctx context.Context, req *pb.FriendPutInHandleReq) (*pb.FriendPutInHandleResp, error) {
	return s.uc.FriendPutInHandle(ctx, req)
}
func (s *SocialService) FriendPutInList(ctx context.Context, req *pb.FriendPutInListReq) (*pb.FriendPutInListResp, error) {
	data, err := s.uc.ListFriendReqByUid(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var list []*pb.FriendRequests
	copier.Copy(&list, &data)

	return &pb.FriendPutInListResp{List: list}, nil
}
func (s *SocialService) FriendList(ctx context.Context, req *pb.FriendListReq) (*pb.FriendListResp, error) {
	data, err := s.uc.ListFriendByUid(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var list []*pb.Friends
	copier.Copy(&list, &data)

	return &pb.FriendListResp{List: list}, nil
}
func (s *SocialService) GroupCreate(ctx context.Context, req *pb.GroupCreateReq) (*pb.GroupCreateResp, error) {
	return s.uc.GroupCreate(ctx, req)
}
func (s *SocialService) GroupPutin(ctx context.Context, req *pb.GroupPutinReq) (*pb.GroupPutinResp, error) {
	return s.uc.GroupPutin(ctx, req)
}
func (s *SocialService) GroupPutinList(ctx context.Context, req *pb.GroupPutinListReq) (*pb.GroupPutinListResp, error) {
	return s.uc.GroupPutinList(ctx, req)
}
func (s *SocialService) GroupPutInHandle(ctx context.Context, req *pb.GroupPutInHandleReq) (*pb.GroupPutInHandleResp, error) {
	return s.uc.GroupPutInHandle(ctx, req)
}
func (s *SocialService) GroupList(ctx context.Context, req *pb.GroupListReq) (*pb.GroupListResp, error) {
	return s.uc.GroupList(ctx, req)
}
func (s *SocialService) GroupUsers(ctx context.Context, req *pb.GroupUsersReq) (*pb.GroupUsersResp, error) {
	return s.uc.GroupUsers(ctx, req)
}
