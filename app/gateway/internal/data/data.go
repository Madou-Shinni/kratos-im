package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/kafka"
	etcdv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
	"kratos-im/api/im"
	"kratos-im/api/social"
	"kratos-im/app/gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGatewayRepo, NewMQClient, NewIMServiceClient, NewRegistrar, NewDiscovery, NewSocialServiceClient)

// Data .
type Data struct {
	kafkaBroker  broker.Broker
	imClient     im.IMClient
	socialClient social.SocialClient
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, kafkaBroker broker.Broker, imClient im.IMClient, socialClient social.SocialClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		kafkaBroker:  kafkaBroker,
		imClient:     imClient,
		socialClient: socialClient,
	}, cleanup, nil
}

func RpcConn(serviceName string, r registry.Discovery) *ggrpc.ClientConn {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+serviceName),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return conn
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	cfg := etcdv3.Config{
		Endpoints: conf.Etcd.Endpoints,
	}
	cli, err := etcdv3.New(cfg)
	if err != nil {
		panic(err)
	}
	return etcd.New(cli)
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	cfg := etcdv3.Config{
		Endpoints: conf.Etcd.Endpoints,
	}
	cli, err := etcdv3.New(cfg)
	if err != nil {
		panic(err)
	}
	return etcd.New(cli)
}

func NewMQClient(c *conf.Data, logger log.Logger) broker.Broker {
	log := log.NewHelper(logger)
	b := kafka.NewBroker(
		broker.WithAddress(c.Kafka.Brokers...),
		broker.WithCodec("json"),
	)

	b.Init()

	if err := b.Connect(); err != nil {
		log.Errorf("cant connect to broker, skip: %v", err)
		return nil
	}

	return b
}

// NewIMServiceClient im服务
func NewIMServiceClient(dis *conf.Discovery, r registry.Discovery) im.IMClient {
	conn := RpcConn(dis.Service.Im, r)
	c := im.NewIMClient(conn)
	return c
}

// NewSocialServiceClient social服务
func NewSocialServiceClient(dis *conf.Discovery, r registry.Discovery) social.SocialClient {
	conn := RpcConn(dis.Service.Social, r)
	c := social.NewSocialClient(conn)
	return c
}
