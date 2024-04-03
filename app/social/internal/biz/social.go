package biz

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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
	ErrorGroupReqAgree   = errorx.ErrorBus("群申请已同意")
	ErrorGroupReqRefuse  = errorx.ErrorBus("群申请已拒绝")
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
	ListGroupByUid(ctx context.Context, uid string) ([]*model.Groups, error)
	SaveGroup(ctx context.Context, data model.Groups) (uint64, error)
	SaveGroupMember(ctx context.Context, data model.GroupMembers) error
	FirstGroupMemberByGidUid(ctx context.Context, gid uint64, uid string) (*model.GroupMembers, error)
	ListGroupMemberByGid(ctx context.Context, gid uint64) ([]*model.GroupMembers, error)
	FirstGroupReqByGidUid(ctx context.Context, gid uint64, uid string) (*model.GroupRequests, error)
	SaveGroupReq(ctx context.Context, data model.GroupRequests) (uint64, error)
	FirstGroupById(ctx context.Context, id uint64) (*model.Groups, error)
	FirstGroupReqById(ctx context.Context, id uint64) (*model.GroupRequests, error)
	UpdateGroupReq(ctx context.Context, freq *model.GroupRequests) error
	ListGroupReqByGid(ctx context.Context, groupId uint64) ([]*model.GroupRequests, error)
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

// GroupCreate 创建群组
func (uc *SocialUsecase) GroupCreate(ctx context.Context, req *pb.GroupCreateReq) (*pb.GroupCreateResp, error) {
	var gid uint64
	var err error
	err = uc.tx.ExecTx(ctx, func(ctx context.Context) error {
		gid, err = uc.repo.SaveGroup(ctx, model.Groups{
			Name:       req.Name,
			Icon:       req.Icon,
			Status:     0,
			CreatorUid: req.CreatorUid,
			GroupType:  0,
			CreatedAt:  time.Now(),
		})
		if err != nil {
			return err
		}

		err = uc.repo.SaveGroupMember(ctx, model.GroupMembers{
			GroupId:   gid,
			UserId:    req.CreatorUid,
			RoleLevel: constants.CreatorGroupRoleLevel,
			JoinTime:  time.Now(),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.GroupCreateResp{Id: gid}, nil
}

// GroupPutin 申请加入群组
func (uc *SocialUsecase) GroupPutin(ctx context.Context, req *pb.GroupPutinReq) (*pb.GroupPutinResp, error) {
	//  1. 普通用户申请 ： 如果群无验证直接进入
	//  2. 群成员邀请： 如果群无验证直接进入
	//  3. 群管理员/群创建者邀请：直接进入群
	var (
		inviteGroupMember *model.GroupMembers
		userGroupMember   *model.GroupMembers
		groupInfo         *model.Groups

		err error
	)

	// 查询群成员
	userGroupMember, err = uc.repo.FirstGroupMemberByGidUid(ctx, req.GroupId, req.ReqId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if userGroupMember != nil {
		return &pb.GroupPutinResp{}, nil
	}

	// 查询入群申请
	groupReq, err := uc.repo.FirstGroupReqByGidUid(ctx, req.GroupId, req.ReqId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if groupReq != nil {
		return &pb.GroupPutinResp{}, nil
	}

	groupReq = &model.GroupRequests{
		ReqId:   req.ReqId,
		GroupId: req.GroupId,
		ReqMsg:  req.ReqMsg,
		ReqTime: sql.NullTime{
			Time:  time.Unix(req.ReqTime, 0),
			Valid: true,
		},
		JoinSource:    int(req.JoinSource),
		InviterUserId: req.InviterUid,
		HandleResult:  constants.HandleResultNone,
	}

	createGroupMember := func() {
		if err != nil {
			return
		}
		err = uc.createGroupMember(ctx, req)
	}

	groupInfo, err = uc.repo.FirstGroupById(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	// 验证是否要验证
	if !groupInfo.IsVerify {
		// 不需要
		defer createGroupMember()

		groupReq.HandleResult = constants.HandleResultAgree

		return uc.createGroupReq(ctx, *groupReq, true)
	}

	// 验证进群方式
	if constants.GroupJoinSource(req.JoinSource) == constants.PutInGroupJoinSource {
		// 申请
		return uc.createGroupReq(ctx, *groupReq, false)
	}

	inviteGroupMember, err = uc.repo.FirstGroupMemberByGidUid(ctx, req.GroupId, req.InviterUid)
	if err != nil {
		return nil, err
	}

	if inviteGroupMember.RoleLevel == constants.CreatorGroupRoleLevel || inviteGroupMember.RoleLevel == constants.ManagerGroupRoleLevel {
		// 是管理者或创建者邀请
		defer createGroupMember()

		groupReq.HandleResult = constants.HandleResultAgree
		groupReq.HandleUserId = req.InviterUid
		return uc.createGroupReq(ctx, *groupReq, true)
	}
	return uc.createGroupReq(ctx, *groupReq, false)
}

func (uc *SocialUsecase) createGroupReq(ctx context.Context, groupReq model.GroupRequests, isPass bool) (*pb.GroupPutinResp, error) {
	_, err := uc.repo.SaveGroupReq(ctx, groupReq)
	if err != nil {
		return nil, err
	}

	if isPass {
		return &pb.GroupPutinResp{GroupId: groupReq.GroupId}, nil
	}

	return &pb.GroupPutinResp{}, nil
}

func (uc *SocialUsecase) createGroupMember(ctx context.Context, req *pb.GroupPutinReq) error {
	groupMember := model.GroupMembers{
		GroupId:     req.GroupId,
		UserId:      req.ReqId,
		RoleLevel:   constants.AtLargeGroupRoleLevel,
		OperatorUid: req.InviterUid,
	}
	err := uc.repo.SaveGroupMember(ctx, groupMember)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SocialUsecase) GroupPutInHandle(ctx context.Context, req *pb.GroupPutInHandleReq) (*pb.GroupPutInHandleResp, error) {
	// 查询群申请记录
	greq, err := uc.repo.FirstGroupReqById(ctx, uint64(req.GroupReqId))
	if err != nil {
		return nil, err
	}
	// 验证是否有处理
	switch greq.HandleResult {
	case constants.HandleResultAgree:
		return nil, ErrorGroupReqAgree
	case constants.HandleResultRefuse:
		return nil, ErrorGroupReqRefuse
	case constants.HandleResultNone:
	default:
	}

	greq.HandleResult = constants.HandleResult(req.HandleResult)
	greq.HandleTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = uc.tx.ExecTx(ctx, func(ctx context.Context) error {
		// 更新群申请记录
		err = uc.repo.UpdateGroupReq(ctx, greq)
		if err != nil {
			return err
		}
		// 同意添加群成员
		if constants.HandleResult(req.HandleResult) == constants.HandleResultAgree {
			err = uc.repo.SaveGroupMember(ctx, model.GroupMembers{
				UserId:      greq.ReqId,
				GroupId:     greq.GroupId,
				RoleLevel:   constants.AtLargeGroupRoleLevel,
				JoinTime:    time.Now(),
				JoinSource:  greq.JoinSource,
				InviterUid:  greq.InviterUserId,
				OperatorUid: greq.HandleUserId,
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

	return &pb.GroupPutInHandleResp{
		GroupId: req.GroupId,
	}, nil
}

func (uc *SocialUsecase) GroupPutinList(ctx context.Context, req *pb.GroupPutinListReq) (*pb.GroupPutinListResp, error) {
	groupReqs, err := uc.repo.ListGroupReqByGid(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	var list []*pb.GroupRequests

	err = copier.Copy(&groupReqs, &groupReqs)
	if err != nil {
		return nil, err
	}

	return &pb.GroupPutinListResp{List: list}, nil
}

func (uc *SocialUsecase) GroupList(ctx context.Context, req *pb.GroupListReq) (*pb.GroupListResp, error) {
	groups, err := uc.repo.ListGroupByUid(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var list []*pb.Groups
	err = copier.Copy(&list, &groups)
	if err != nil {
		return nil, err
	}

	return &pb.GroupListResp{List: list}, nil
}

func (uc *SocialUsecase) GroupUsers(ctx context.Context, req *pb.GroupUsersReq) (*pb.GroupUsersResp, error) {
	groupMembers, err := uc.repo.ListGroupMemberByGid(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}

	var list = make([]*pb.GroupMembers, 0, len(groupMembers))

	err = copier.Copy(&list, &groupMembers)
	if err != nil {
		return nil, err
	}

	return &pb.GroupUsersResp{
		List: list,
	}, nil
}
