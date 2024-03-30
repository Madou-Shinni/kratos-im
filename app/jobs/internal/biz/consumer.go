package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-im/constants"
	"kratos-im/model"
	"kratos-im/pkg/rws"
)

// ConsumerRepo is a Greater repo.
type ConsumerRepo interface {
	Save(ctx context.Context, chatLog model.ChatLog) error
	UpdateMsg(ctx context.Context, chatLog *model.ChatLog) error
}

// ConsumerUsecase is a Consumer usecase.
type ConsumerUsecase struct {
	repo     ConsumerRepo
	log      *log.Helper
	wsClient rws.IClient
}

// NewConsumerUsecase new a Consumer usecase.
func NewConsumerUsecase(repo ConsumerRepo, logger log.Logger, wsClient rws.IClient) *ConsumerUsecase {
	return &ConsumerUsecase{repo: repo, log: log.NewHelper(logger), wsClient: wsClient}
}

func (u *ConsumerUsecase) HandleMsgTransfer(ctx context.Context, topic string, headers broker.Headers, msg *rws.MsgChatTransfer) error {
	chatLog := model.ChatLog{
		ConversationId: msg.ConversationId,
		SendId:         msg.SendId,
		RecvId:         msg.RecvId,
		MsgFrom:        0,
		ChatType:       msg.ChatType,
		MsgType:        msg.MType,
		MsgContent:     msg.Content,
		SendTime:       msg.SendTime,
		Status:         0,
	}

	// 保存数据
	err := u.repo.Save(ctx, chatLog)
	if err != nil {
		u.log.Errorf("HandleMsgTransfer Save: %v", err)
		return err
	}

	// 更新会话
	err = u.repo.UpdateMsg(ctx, &chatLog)
	if err != nil {
		u.log.Errorf("HandleMsgTransfer UpdateMsg: %v", err)
		return err
	}

	// 推送消息(推送给 ws server)
	switch chatLog.ChatType {
	case constants.ChatTypeSingle:
		return u.sendSingle(msg)
	case constants.ChatTypeGroup:
		return u.sendGroup()
	}

	return nil
}

func (u *ConsumerUsecase) sendSingle(data *rws.MsgChatTransfer) error {
	return u.wsClient.Send(rws.Message{
		FrameType: rws.FrameData,
		Method:    "push",
		FromId:    constants.SystemRootUid,
		//ToId:      "",
		Data: data,
	})
}

func (u *ConsumerUsecase) sendGroup() error {
	// TODO: 查询群用户
	return nil
}
