package data

import (
	"context"
	"kratos-im/app/im/internal/biz"
	"kratos-im/model"

	"github.com/go-kratos/kratos/v2/log"
)

type imRepo struct {
	data *Data
	log  *log.Helper
}

// NewIMRepo .
func NewIMRepo(data *Data, logger log.Logger) biz.IMRepo {
	return &imRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *imRepo) Save(ctx context.Context, chatLog model.ChatLog) error {
	_, err := r.data.mongoDatabase.Collection(chatLog.Collection()).InsertOne(ctx, chatLog)
	if err != nil {
		return err
	}
	return nil
}
