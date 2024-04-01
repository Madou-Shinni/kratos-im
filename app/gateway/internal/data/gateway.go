package data

import (
	"context"
	imPb "kratos-im/api/im"
	"kratos-im/model"

	"kratos-im/app/gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type gatewayRepo struct {
	data *Data
	log  *log.Helper
}

// NewGatewayRepo .
func NewGatewayRepo(data *Data, logger log.Logger) biz.GatewayRepo {
	return &gatewayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *gatewayRepo) Save(ctx context.Context, data model.ChatLog) error {
	_, err := r.data.imClient.CreateChatLog(ctx, &imPb.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		MsgType:        int32(data.MsgType),
		MsgContent:     data.MsgContent,
		ChatType:       int32(data.ChatType),
		SendTime:       data.SendTime,
	})
	if err != nil {
		return err
	}
	return nil
}
