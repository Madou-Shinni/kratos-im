package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-im/app/im/internal/biz"
	"kratos-im/pkg/rws"
	"kratos-im/pkg/tools"
	"time"
)

// IMService is a greeter service.
type IMService struct {
	uc          *biz.IMUsecase
	KafkaBroker broker.Broker
}

// NewIMService new a greeter service.
func NewIMService(uc *biz.IMUsecase, kafkaBroker broker.Broker) *IMService {
	return &IMService{
		uc:          uc,
		KafkaBroker: kafkaBroker,
	}
}

// CreateChatLog 创建私聊消息
func (s *IMService) CreateChatLog(ctx context.Context, data *rws.Chat, userID string) error {
	if data.ConversationId == "" {
		data.ConversationId = tools.CombineId(userID, data.RecvId)
	}

	data.SendId = userID
	data.SendTime = time.Now().UnixMilli()

	return s.uc.CreateChatLog(ctx, data)
}
