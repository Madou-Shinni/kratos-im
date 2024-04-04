package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-im/constants"
	"kratos-im/pkg/rws"
	"strconv"
)

type BaseMsgTransfer struct {
	repo     ConsumerRepo
	log      *log.Helper
	wsClient rws.IClient
}

// NewBaseMsgTransfer new a Consumer usecase.
func NewBaseMsgTransfer(repo ConsumerRepo, logger log.Logger, wsClient rws.IClient) *BaseMsgTransfer {
	return &BaseMsgTransfer{repo: repo, log: log.NewHelper(logger), wsClient: wsClient}
}

func (r *BaseMsgTransfer) Transfer(ctx context.Context, data *rws.Push) error {
	switch data.ChatType {
	case constants.ChatTypeGroup:
		return r.sendGroup(data) // 群聊
	case constants.ChatTypeSingle:
		return r.sendSingle(data) // 单聊
	}

	return nil
}

func (r *BaseMsgTransfer) sendSingle(data *rws.Push) error {
	return r.wsClient.Send(rws.Message{
		FrameType: rws.FrameData,
		Method:    "push",
		FromId:    constants.SystemRootUid,
		Data:      data,
	})
}

func (r *BaseMsgTransfer) sendGroup(data *rws.Push) error {
	// 查询群用户
	gid, err := strconv.ParseUint(data.RecvId, 10, 64)
	if err != nil {
		return err
	}
	members, err := r.repo.ListGroupMembersByGid(context.Background(), gid)
	if err != nil {
		return err
	}

	var uids = make([]string, 0, len(members))
	for _, v := range members {
		if v.UserId == data.SendId {
			// 不推送给自己
			continue
		}
		uids = append(uids, v.UserId)
	}

	data.RecvIds = uids

	return r.wsClient.Send(rws.Message{
		FrameType: rws.FrameData,
		Method:    "push",
		FromId:    constants.SystemRootUid,
		Data:      data,
	})
}
