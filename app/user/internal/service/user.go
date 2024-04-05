package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
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

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	data, err := s.uc.Login(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: data.Token,
		UserInfo: &pb.LoginReply_UserInfo{
			Token:    data.Token,
			UserId:   data.ID,
			Avatar:   data.Avatar,
			Nickname: data.Nickname,
		},
	}, nil
}

func (s *UserService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResp, error) {
	return s.uc.ListByIds(ctx, req.Ids)
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
