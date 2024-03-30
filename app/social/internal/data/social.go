package data

import (
	"context"
	"kratos-im/model"

	"kratos-im/app/social/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type socialRepo struct {
	data *Data
	log  *log.Helper
}

// NewSocialRepo .
func NewSocialRepo(data *Data, logger log.Logger) biz.SocialRepo {
	return &socialRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// SaveFriendReq 保存好友申请记录
func (r *socialRepo) SaveFriendReq(ctx context.Context, data *model.FriendRequests) error {
	return r.data.DB(ctx).Create(&data).Error
}

func (r *socialRepo) SaveFriend(ctx context.Context, g *biz.Social) (*biz.Social, error) {
	return g, nil
}

// UpdateFriendReq 更新好友申请记录
func (r *socialRepo) UpdateFriendReq(ctx context.Context, data *model.FriendRequests) error {
	err := r.data.DB(ctx).Model(&model.FriendRequests{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// SaveFriends 添加好友
func (r *socialRepo) SaveFriends(ctx context.Context, data ...*model.Friends) error {
	err := r.data.DB(ctx).Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *socialRepo) FindByID(context.Context, int64) (*biz.Social, error) {
	return nil, nil
}

func (r *socialRepo) ListByHello(context.Context, string) ([]*biz.Social, error) {
	return nil, nil
}

// ListFriendByUid 好友列表
func (r *socialRepo) ListFriendByUid(ctx context.Context, uid string) ([]*model.Friends, error) {
	var data []*model.Friends
	err := r.data.db.Find(&data, "user_id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// FirstFriendByUidFid 通过fid和uid查找好友
func (r *socialRepo) FirstFriendByUidFid(ctx context.Context, uid, fid string) (*model.Friends, error) {
	var data model.Friends
	err := r.data.db.First(&data, "user_id = ? AND friend_uid = ?", uid, fid).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FirstFriendReqByRuidUid 通过ruid和uid查找请求记录
func (r *socialRepo) FirstFriendReqByRuidUid(ctx context.Context, ruid, uid string) (*model.FriendRequests, error) {
	var data model.FriendRequests
	err := r.data.db.First(&data, "req_uid = ? AND user_id = ?", ruid, uid).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FirstFriendReqByRid 通过id查找请求记录
func (r *socialRepo) FirstFriendReqByRid(ctx context.Context, rid uint64) (*model.FriendRequests, error) {
	var data model.FriendRequests
	err := r.data.db.First(&data, "id = ?", rid).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// ListFriendReqByUid 通过uid查询好友申请记录
func (r *socialRepo) ListFriendReqByUid(ctx context.Context, uid string) ([]*model.FriendRequests, error) {
	var data []*model.FriendRequests
	err := r.data.db.Find(&data, "user_id = ? or req_uid = ?", uid, uid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
