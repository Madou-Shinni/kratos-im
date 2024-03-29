package data

import (
	"context"
	"kratos-im/model"

	"kratos-im/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) Save(ctx context.Context, user *model.User) (*model.User, error) {
	err := u.data.db.Where("github_id = ?", user.GithubId).Attrs(model.User{
		ID:       user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Sex:      user.Sex,
		GithubId: user.GithubId,
	}).FirstOrCreate(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
