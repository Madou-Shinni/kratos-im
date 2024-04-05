package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	v1 "kratos-im/api/user"
	"kratos-im/app/user/internal/conf"
	"kratos-im/constants"
	"kratos-im/model"
	"kratos-im/pkg/tools"
	"time"
)

// LoginResp is a login response.
type LoginResp struct {
	Token string
	*model.User
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *model.User) (*model.User, error)
	ListByIds(ctx context.Context, ids []string) ([]*model.User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo  UserRepo
	log   *log.Helper
	c     *conf.Oauth2
	cauth *conf.Auth
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger, c *conf.Oauth2, cauth *conf.Auth) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger), c: c, cauth: cauth}
}

func (uc *UserUsecase) Login(ctx context.Context, code string) (*LoginResp, error) {
	// 获取github配置
	clientId := uc.c.Github.ClientId
	clientSecret := uc.c.Github.ClientSecret

	// 获取github token
	result, err := tools.LoginGithub(tools.LoginGithubReq{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Code:         code,
	})
	if err != nil {
		uc.log.Errorf("tools.LoginGithub(tools.LoginGithubReq{}) err: %v", err)
		return nil, err
	}

	// 获取用户信息
	info, err := result.GetUserInfo()
	if err != nil {
		uc.log.Errorf("result.GetUserInfo() err: %v", err)
		return nil, err
	}

	// 保存用户信息
	user := model.User{
		ID:       uuid.NewString(),
		Nickname: info.Name,
		Avatar:   info.Avatar,
		Email:    info.Email,
		Sex:      0,
		GithubId: info.ID,
	}

	userInfo, err := uc.repo.Save(ctx, &user)
	if err != nil {
		uc.log.Errorf("uc.repo.Save(ctx, &user) err: %v", err)
		return nil, err
	}

	// 生成token
	mapClaims := jwt.MapClaims{
		constants.CtxUserIDKey: userInfo.ID,
	}
	expire := time.Duration(uc.cauth.Expire) * time.Second
	secret := uc.cauth.Key
	token, err := tools.GenToken(mapClaims, expire, secret)
	if err != nil {
		uc.log.Errorf("tools.GenToken(mapClaims, expire, secret) err: %v", err)
		return nil, err
	}

	return &LoginResp{
		Token: token,
		User:  userInfo,
	}, nil
}

func (uc *UserUsecase) ListByIds(ctx context.Context, ids []string) (*v1.ListResp, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	users, err := uc.repo.ListByIds(ctx, ids)
	if err != nil {
		uc.log.Errorf("uc.repo.ListByIds(ctx, ids) err: %v", err)
		return nil, err
	}

	var userMap = make(map[string]*v1.ListResp_UserInfo, len(users))

	for _, user := range users {
		userMap[user.ID] = &v1.ListResp_UserInfo{
			UserId:   user.ID,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		}
	}

	return &v1.ListResp{
		Users: userMap,
	}, nil
}
