package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	pb "kratos-im/api/im"
	v1 "kratos-im/api/im"
	"kratos-im/app/im/internal/biz"
	"kratos-im/pkg/rws"
	"kratos-im/pkg/tools"
	"time"
)

// IMService is a greeter service.
type IMService struct {
	v1.UnimplementedIMServer
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

func (s *IMService) GetChatLog(ctx context.Context, req *pb.GetChatLogReq) (*pb.GetChatLogResp, error) {
	list, err := s.uc.GetChatLog(ctx, req)
	if err != nil {
		return nil, nil
	}

	data := make([]*pb.ChatLog, 0, len(list))

	for _, v := range list {
		data = append(data, &pb.ChatLog{
			Id:             v.ID.Hex(),
			ConversationId: v.ConversationId,
			SendId:         v.SendId,
			RecvId:         v.RecvId,
			MsgType:        int32(v.MsgType),
			MsgContent:     v.MsgContent,
			ChatType:       int32(v.ChatType),
			SendTime:       v.SendTime,
			ReadRecords:    nil,
		})
	}

	return &pb.GetChatLogResp{
		List: data,
	}, nil
}
func (s *IMService) SetUpUserConversation(ctx context.Context, req *pb.SetUpUserConversationReq) (*pb.SetUpUserConversationResp, error) {
	return s.uc.SetUpUserConversation(ctx, req)
}
func (s *IMService) GetConversations(ctx context.Context, req *pb.GetConversationsReq) (*pb.GetConversationsResp, error) {
	return s.uc.GetConversations(ctx, req)
}
func (s *IMService) PutConversations(ctx context.Context, req *pb.PutConversationsReq) (*pb.PutConversationsResp, error) {
	return s.uc.PutConversations(ctx, req)
}
func (s *IMService) CreateGroupConversation(ctx context.Context, req *pb.CreateGroupConversationReq) (*pb.CreateGroupConversationResp, error) {
	return &pb.CreateGroupConversationResp{}, nil
}
