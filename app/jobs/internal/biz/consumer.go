package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-im/model"
	"kratos-im/pkg/rws"
)

// ConsumerRepo is a Greater repo.
type ConsumerRepo interface {
	Save(ctx context.Context, chatLog model.ChatLog) error
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

func (u ConsumerUsecase) HandleMsgTransfer(ctx context.Context, topic string, headers broker.Headers, msg *rws.MsgChatTransfer) error {
	// 保存数据
	err := u.repo.Save(ctx, model.ChatLog{
		ConversationId: msg.ConversationId,
		SendId:         msg.SendId,
		RecvId:         msg.RecvId,
		MsgFrom:        0,
		MsgType:        msg.MType,
		MsgContent:     msg.Content,
		SendTime:       msg.SendTime,
		Status:         0,
	})
	if err != nil {
		u.log.Errorf("HandleMsgTransfer Save: %v", err)
		return err
	}

	// 推送消息(推送给 ws server)
	err = u.wsClient.Send(rws.Message{
		FrameType: rws.FrameData,
		Method:    "push",
		//FromId:    constants.SystemRootUid,
		//ToId:      "",
		Data: *msg,
	})
	if err != nil {
		u.log.Errorf("HandleMsgTransfer Send: %v", err)
		return err
	}

	return nil
}
