package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-im/pkg/rws"

	"kratos-im/app/jobs/internal/biz"
)

// ConsumerService is a greeter service.
type ConsumerService struct {
	uc *biz.ConsumerUsecase
}

// NewConsumerService new a greeter service.
func NewConsumerService(uc *biz.ConsumerUsecase) *ConsumerService {
	return &ConsumerService{uc: uc}
}

// HandleMsgTransfer implements the ConsumerService.
func (s *ConsumerService) HandleMsgTransfer(ctx context.Context, topic string, headers broker.Headers, msg *rws.MsgChatTransfer) error {
	return s.uc.HandleMsgTransfer(ctx, topic, headers, msg)
}

// HandleMsgReadTransfer implements the ConsumerService.
func (s *ConsumerService) HandleMsgReadTransfer(ctx context.Context, topic string, headers broker.Headers, msg *rws.MsgMarkReadTransfer) error {
	return s.uc.HandleMsgReadTransfer(ctx, topic, headers, msg)
}
