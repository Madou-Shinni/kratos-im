package biz

import (
	"context"
	"kratos-im/model"
	"kratos-im/pkg/rws"

	"github.com/go-kratos/kratos/v2/log"
)

// IM is a IM model.
type IM struct {
	Hello string
}

// IMRepo is a Greater repo.
type IMRepo interface {
	Save(ctx context.Context, chatLog model.ChatLog) error
}

// IMUsecase is a IM usecase.
type IMUsecase struct {
	repo IMRepo
	log  *log.Helper
}

// NewIMUsecase new a IM usecase.
func NewIMUsecase(repo IMRepo, logger log.Logger) *IMUsecase {
	return &IMUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *IMUsecase) CreateChatLog(ctx context.Context, data *rws.Chat) error {
	chatLog := model.ChatLog{
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ConversationId: data.ConversationId,
		MsgFrom:        0,
		MsgType:        data.Msg.MType,
		MsgContent:     data.Msg.Content,
		SendTime:       data.SendTime,
		Status:         0,
	}
	return uc.repo.Save(ctx, chatLog)
}
