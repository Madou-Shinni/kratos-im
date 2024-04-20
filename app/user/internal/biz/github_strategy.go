package biz

import (
	"context"
	"github.com/google/uuid"
	"kratos-im/api/errorx"
	v1 "kratos-im/api/user"
	"kratos-im/app/user/internal/conf"
	"kratos-im/model"
	"kratos-im/pkg/tools"
)

var ErrorGithubAuthFailed = errorx.ErrorBus("github login failed")
var ErrorGithubUserProfileFailed = errorx.ErrorBus("github login get user profile failed")

type GithubLogin struct {
	repo UserRepo
	c    *conf.Oauth2
}

func NewGithubLogin(repo UserRepo, oauth2 *conf.Oauth2) *GithubLogin {
	return &GithubLogin{repo: repo, c: oauth2}
}

func (s *GithubLogin) Login(ctx context.Context, req *v1.LoginRequest) (*model.User, error) {
	acReq := req.Payload.(*v1.LoginRequest_Github).Github

	// 获取github配置
	clientId := s.c.Github.ClientId
	clientSecret := s.c.Github.ClientSecret

	// 获取github token
	result, err := tools.LoginGithub(tools.LoginGithubReq{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Code:         acReq.Code,
	})
	if err != nil {
		return nil, err
	}
	if result.AccessToken == "" {
		return nil, ErrorGithubAuthFailed
	}

	// 获取用户信息
	info, err := result.GetUserInfo()
	if err != nil {
		return nil, ErrorGithubUserProfileFailed
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

	userInfo, err := s.repo.Save(ctx, &user)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
