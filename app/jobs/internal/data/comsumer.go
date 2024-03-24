package data

import (
	"context"

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

func (r *consumerRepo) Save(ctx context.Context, g *biz.Consumer) (*biz.Consumer, error) {
	return g, nil
}

func (r *consumerRepo) Update(ctx context.Context, g *biz.Consumer) (*biz.Consumer, error) {
	return g, nil
}

func (r *consumerRepo) FindByID(context.Context, int64) (*biz.Consumer, error) {
	return nil, nil
}

func (r *consumerRepo) ListByHello(context.Context, string) ([]*biz.Consumer, error) {
	return nil, nil
}

func (r *consumerRepo) ListAll(context.Context) ([]*biz.Consumer, error) {
	return nil, nil
}
