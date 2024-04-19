package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kratos-im/api/errorx"
	v1 "kratos-im/api/user"
	"kratos-im/app/user/internal/conf"
	"kratos-im/constants"
	"kratos-im/model"
	"kratos-im/pkg/tools"
	"time"
)

var ErrorNotSupportLoginType = errorx.ErrorBus("not support login type")
var ErrorUserExist = errorx.ErrorBus("user exist")

type ILoginStrategy interface {
	Login(ctx context.Context, req *v1.LoginRequest) (*model.User, error)
}

// LoginResp is a login response.
type LoginResp struct {
	Token string
	*model.User
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *model.User) (*model.User, error)
	ListByIds(ctx context.Context, ids []string) ([]*model.User, error)
	FirstByAccountOrEmail(ctx context.Context, account, email string) (*model.User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo       UserRepo
	lgstrategy map[constants.LoginType]ILoginStrategy
	log        *log.Helper
	c          *conf.Oauth2
	cauth      *conf.Auth
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger, c *conf.Oauth2, cauth *conf.Auth) *UserUsecase {
	uc := &UserUsecase{repo: repo, log: log.NewHelper(logger), c: c, cauth: cauth}

	uc.lgstrategy = make(map[constants.LoginType]ILoginStrategy)
	uc.lgstrategy[constants.LoginTypeAccount] = NewAccountLogin(repo)
	uc.lgstrategy[constants.LoginTypeGithub] = NewGithubLogin(repo)

	return uc
}

func (uc *UserUsecase) Login(ctx context.Context, req *v1.LoginRequest) (*LoginResp, error) {
	var userInfo *model.User
	var err error

	if _, ok := uc.lgstrategy[constants.LoginType(req.Type)]; !ok {
		return nil, ErrorNotSupportLoginType
	}

	userInfo, err = uc.lgstrategy[constants.LoginType(req.Type)].Login(ctx, req)
	if err != nil {
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

func (uc *UserUsecase) Register(ctx context.Context, account string, password string) error {
	// 查询用户是否存在
	user, err := uc.repo.FirstByAccountOrEmail(ctx, account, "")
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if user != nil {
		return ErrorUserExist
	}

	// 保存用户信息
	u := &model.User{
		ID:       uuid.NewString(),
		Account:  account,
		Password: tools.BcryptHash(password),
	}

	_, err = uc.repo.Save(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
