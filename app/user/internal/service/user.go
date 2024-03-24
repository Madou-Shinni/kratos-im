package service

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	v1 "kratos-im/api/user"
	"kratos-im/app/user/internal/biz"
	"kratos-im/app/user/internal/conf"
	"kratos-im/constants"
	"kratos-im/pkg/tools"
	"time"

	pb "kratos-im/api/user"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	c   *conf.Auth
	rdb redis.Cmdable
}

func NewUserService(uc *biz.UserUsecase, c *conf.Auth, rdb redis.Cmdable) *UserService {
	return &UserService{
		uc:  uc,
		c:   c,
		rdb: rdb,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *UserService) SetRootToken() error {
	ctx := context.Background()
	// 生成jwt token
	exp := time.Hour * 24 * 365
	token, err := tools.GenToken(jwt.MapClaims{
		constants.CtxUserIDKey: constants.SystemRootUid,
	}, exp, s.c.Key)
	if err != nil {
		return err
	}
	// 保存到redis
	err = s.rdb.Set(ctx, constants.SystemRootUid, token, -1).Err()
	if err != nil {
		return err
	}

	return nil
}
