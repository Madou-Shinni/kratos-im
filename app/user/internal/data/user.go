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
	var userInfo model.User
	err := u.data.db.Where("github_id = ?", user.GithubId).Attrs(model.User{
		ID:       user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Sex:      user.Sex,
		GithubId: user.GithubId,
	}).FirstOrCreate(&userInfo).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepo) ListByIds(ctx context.Context, ids []string) ([]*model.User, error) {
	var users []*model.User
	err := u.data.db.Where("id in ?", ids).First(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepo) FirstByAccountOrEmail(ctx context.Context, account, email string) (*model.User, error) {
	var user model.User
	if account != "" {
		err := u.data.db.Where("account = ?", account).First(&user).Error
		if err != nil {
			return nil, err
		}
	}

	if email != "" {
		err := u.data.db.Where("email = ?", email).First(&user).Error
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
