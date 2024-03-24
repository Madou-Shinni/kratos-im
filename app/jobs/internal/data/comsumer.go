package data

import (
	"context"
	"kratos-im/model"

	"kratos-im/app/jobs/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type consumerRepo struct {
	data *Data
	log  *log.Helper
}

// NewConsumerRepo .
func NewConsumerRepo(data *Data, logger log.Logger) biz.ConsumerRepo {
	return &consumerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *consumerRepo) Save(ctx context.Context, chatLog model.ChatLog) error {
	_, err := r.data.mongoDatabase.Collection(chatLog.Collection()).InsertOne(ctx, chatLog)
	if err != nil {
		return err
	}
	return nil
}
