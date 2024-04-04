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

// FirstGroupById 通过id查找群
func (r *socialRepo) FirstGroupById(ctx context.Context, id uint64) (*model.Groups, error) {
	var data model.Groups
	err := r.data.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
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
	err := r.data.db.Find(&data, "user_id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListGroupByUid 通过uid查询群列表
func (r *socialRepo) ListGroupByUid(ctx context.Context, uid string) ([]*model.Groups, error) {
	var data []*model.Groups
	err := r.data.db.Find(&data, "creator_uid = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SaveGroupReq 创建入群申请
func (r *socialRepo) SaveGroupReq(ctx context.Context, data model.GroupRequests) (uint64, error) {
	return data.ID, r.data.DB(ctx).Create(&data).Error
}

// SaveGroup 创建群
func (r *socialRepo) SaveGroup(ctx context.Context, data model.Groups) (uint64, error) {
	return data.ID, r.data.DB(ctx).Create(&data).Error
}

// SaveGroupMember 添加群成员
func (r *socialRepo) SaveGroupMember(ctx context.Context, data model.GroupMembers) error {
	return r.data.DB(ctx).Create(&data).Error
}

// ListGroupMemberByGid 通过gid查询群成员列表
func (r *socialRepo) ListGroupMemberByGid(ctx context.Context, gid uint64) ([]*model.GroupMembers, error) {
	var data []*model.GroupMembers
	err := r.data.db.Find(&data, "group_id = ?", gid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// FirstGroupMemberByGidUid 通过gid,uid查询群成员
func (r *socialRepo) FirstGroupMemberByGidUid(ctx context.Context, gid uint64, uid string) (*model.GroupMembers, error) {
	var data *model.GroupMembers
	err := r.data.db.First(&data, "group_id = ? AND user_id = ?", gid, uid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// FirstGroupReqByGidUid 通过gid,uid查询入群申请
func (r *socialRepo) FirstGroupReqByGidUid(ctx context.Context, gid uint64, uid string) (*model.GroupRequests, error) {
	var data model.GroupRequests
	err := r.data.db.First(&data, "group_id = ? AND req_id = ?", gid, uid).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FirstGroupReqById 通过id查询入群申请
func (r *socialRepo) FirstGroupReqById(ctx context.Context, id uint64) (*model.GroupRequests, error) {
	var data *model.GroupRequests
	err := r.data.db.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateGroupReq 更新群申请
func (r *socialRepo) UpdateGroupReq(ctx context.Context, freq *model.GroupRequests) error {
	err := r.data.db.Model(&model.GroupRequests{}).Where("id = ?", freq.ID).Save(freq).Error
	if err != nil {
		return err
	}
	return nil
}

// ListGroupReqByGid 通过gid查询入群申请
func (r *socialRepo) ListGroupReqByGid(ctx context.Context, gid uint64) ([]*model.GroupRequests, error) {
	var data []*model.GroupRequests
	err := r.data.db.Find(&data, "group_id = ?", gid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
