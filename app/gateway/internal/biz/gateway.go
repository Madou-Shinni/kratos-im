package biz

import (
	"context"
	v1 "kratos-im/api/gateway"
	"kratos-im/common"
	"kratos-im/model"
	"kratos-im/pkg/rws"

	"github.com/go-kratos/kratos/v2/log"
)

// Gateway is a Gateway model.
type Gateway struct {
	Hello string
}

// GatewayRepo is a Greater repo.
type GatewayRepo interface {
	Save(ctx context.Context, chatLog model.ChatLog) error
}

// GatewayUsecase is a Gateway usecase.
type GatewayUsecase struct {
	repo GatewayRepo
	log  *log.Helper
}

// NewGatewayUsecase new a Gateway usecase.
func NewGatewayUsecase(repo GatewayRepo, logger log.Logger) *GatewayUsecase {
	return &GatewayUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GatewayUsecase) CreateChatLog(ctx context.Context, data *rws.Chat) error {
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

func (uc *GatewayUsecase) GroupPutin(ctx context.Context, req *v1.GroupPutinReq) (*v1.GroupPutinResp, error) {
	// 获取uid
	uid, err := common.GetUidFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	uc.log.Info("GroupPutin uid: %s", uid)

	return &v1.GroupPutinResp{}, nil
}
