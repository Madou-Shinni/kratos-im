package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"kratos-im/app/jobs/internal/conf"
	"kratos-im/app/jobs/internal/service"
	"kratos-im/constants"
)

func NewKafkaServer(c *conf.Data, logger log.Logger, s *service.ConsumerService) *kafka.Server {
	ctx := context.Background()
	server := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Brokers),
		kafka.WithCodec("json"),
	)

	logHelper := log.NewHelper(log.With(logger, "module", "Consumer/server/consumer-service"))

	register(server, s, ctx, logHelper)

	return server
}

func register(svr *kafka.Server, svc *service.ConsumerService, ctx context.Context, log *log.Helper) {
	var err error
	err = kafka.RegisterSubscriber(svr, ctx, constants.TopicMsgTransfer, "kafka", false, svc.HandleMsgTransfer)
	if err != nil {
		goto Err
	}

Err:
	log.Errorf("register subscriber error: %v", err)
	return
}
