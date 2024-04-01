package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	pb "kratos-im/api/gateway"
	"kratos-im/app/gateway/internal/biz"
)

// GatewayService is a greeter service.
type GatewayService struct {
	//v1.UnimplementedGatewayServer

	uc          *biz.GatewayUsecase
	KafkaBroker broker.Broker
}

// NewGatewayService new a greeter service.
func NewGatewayService(uc *biz.GatewayUsecase) *GatewayService {
	return &GatewayService{uc: uc}
}

func (s *GatewayService) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello " + in.Name}, nil
}
