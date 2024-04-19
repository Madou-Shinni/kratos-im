package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"kratos-im/api/errorx"
	v1 "kratos-im/api/user"
	"kratos-im/model"
	"kratos-im/pkg/tools"
)

var ErrorAccountOrEmailNotFound = errorx.ErrorBus("account or email not found")
var ErrorPasswordError = errorx.ErrorBus("password error")

type AccountLogin struct {
	repo UserRepo
}

func NewAccountLogin(repo UserRepo) *AccountLogin {
	return &AccountLogin{repo: repo}
}

func (s *AccountLogin) Login(ctx context.Context, req *v1.LoginRequest) (*model.User, error) {
	acReq := req.Payload.(*v1.LoginRequest_Account).Account

	u, err := s.repo.FirstByAccountOrEmail(ctx, acReq.Account, acReq.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if u == nil {
		return nil, ErrorAccountOrEmailNotFound
	}

	if !tools.BcryptCheck(acReq.Password, u.Password) {
		return nil, ErrorPasswordError
	}

	return u, nil
}
