package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
)

// Consumer is a Consumer model.
type Consumer struct {
	Hello string
}

// ConsumerRepo is a Greater repo.
type ConsumerRepo interface {
}

// ConsumerUsecase is a Consumer usecase.
type ConsumerUsecase struct {
	repo ConsumerRepo
	log  *log.Helper
}

// NewConsumerUsecase new a Consumer usecase.
func NewConsumerUsecase(repo ConsumerRepo, logger log.Logger) *ConsumerUsecase {
	return &ConsumerUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u ConsumerUsecase) HandleMsgTransfer(ctx context.Context, topic string, headers broker.Headers, msg *string) error {
	u.log.Infof("HandleMsgTransfer: %v", *msg)
	return nil
}
