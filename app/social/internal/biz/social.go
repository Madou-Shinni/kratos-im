package biz

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-im/api/errorx"
	pb "kratos-im/api/social"
	"kratos-im/constants"
	"kratos-im/model"
	"time"
)

var (
	ErrorFriendReqRefuse = errorx.ErrorBus("好友申请已拒绝")
	ErrorFriendReqAgree  = errorx.ErrorBus("好友申请已同意")
)

// Social is a Social model.
type Social struct {
	Hello string
}

// SocialRepo is a Greater repo.
type SocialRepo interface {
	FirstFriendByUidFid(ctx context.Context, uid, fid string) (*model.Friends, error)
	FirstFriendReqByRuidUid(ctx context.Context, ruid, uid string) (*model.FriendRequests, error)
	FirstFriendReqByRid(ctx context.Context, rid uint64) (*model.FriendRequests, error)
	SaveFriendReq(ctx context.Context, data *model.FriendRequests) error
	UpdateFriendReq(ctx context.Context, data *model.FriendRequests) error
	SaveFriends(ctx context.Context, data ...*model.Friends) error
	ListFriendByUid(ctx context.Context, uid string) ([]*model.Friends, error)
	ListFriendReqByUid(ctx context.Context, uid string) ([]*model.FriendRequests, error)
}

// SocialUsecase is a Social usecase.
type SocialUsecase struct {
	repo SocialRepo
	log  *log.Helper
	tx   Transaction
}

// NewSocialUsecase new a Social usecase.
func NewSocialUsecase(repo SocialRepo, logger log.Logger, tx Transaction) *SocialUsecase {
	return &SocialUsecase{repo: repo, log: log.NewHelper(logger), tx: tx}
}

// FriendPutIn 好友申请
func (uc *SocialUsecase) FriendPutIn(ctx context.Context, req *pb.FriendPutInReq) (*pb.FriendPutInResp, error) {
	// 通过fid和uid查找好友
	friend, err := uc.repo.FirstFriendByUidFid(ctx, req.UserId, req.ReqUid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if friend != nil {
		return &pb.FriendPutInResp{}, nil
	}

	// 通过Ruid Uid申请记录
	freq, err := uc.repo.FirstFriendReqByRuidUid(ctx, req.ReqUid, req.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if freq != nil {
		return &pb.FriendPutInResp{}, nil
	}

	// 保存好友申请记录
	err = uc.repo.SaveFriendReq(ctx, &model.FriendRequests{
		UserId:       req.UserId,
		ReqUid:       req.ReqUid,
		ReqMsg:       req.ReqMsg,
		ReqTime:      time.Unix(req.ReqTime, 0),
		HandleResult: constants.HandleResultNone,
	})
	if err != nil {
		return nil, err
	}

	return &pb.FriendPutInResp{}, nil
}

// FriendPutInHandle 好友申请处理
func (uc *SocialUsecase) FriendPutInHandle(ctx context.Context, req *pb.FriendPutInHandleReq) (*pb.FriendPutInHandleResp, error) {
	// 查询好友申请记录
	freq, err := uc.repo.FirstFriendReqByRid(ctx, uint64(req.FriendReqId))
	if err != nil {
		return nil, err
	}
	// 验证是否有处理
	switch freq.HandleResult {
	case constants.HandleResultAgree:
		return nil, ErrorFriendReqAgree
	case constants.HandleResultRefuse:
		return nil, ErrorFriendReqRefuse
	case constants.HandleResultNone:
	default:
	}

	freq.HandleResult = constants.HandleResult(req.HandleResult)
	freq.HandledAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = uc.tx.ExecTx(ctx, func(ctx context.Context) error {
		// 更新好友申请记录
		err = uc.repo.UpdateFriendReq(ctx, freq)
		if err != nil {
			return err
		}
		// 同意添加好友
		if constants.HandleResult(req.HandleResult) == constants.HandleResultAgree {
			err = uc.repo.SaveFriends(ctx, &model.Friends{
				UserId:    freq.UserId,
				FriendUid: freq.ReqUid,
				CreatedAt: time.Now(),
			}, &model.Friends{
				UserId:    freq.ReqUid,
				FriendUid: freq.UserId,
				CreatedAt: time.Now(),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.FriendPutInHandleResp{}, nil
}

// ListFriendByUid 好友列表
func (uc *SocialUsecase) ListFriendByUid(ctx context.Context, uid string) ([]*model.Friends, error) {
	data, err := uc.repo.ListFriendByUid(ctx, uid)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ListFriendReqByUid 申请列表
func (uc *SocialUsecase) ListFriendReqByUid(ctx context.Context, uid string) ([]*model.FriendRequests, error) {
	return uc.repo.ListFriendReqByUid(ctx, uid)
}
