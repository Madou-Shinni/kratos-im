package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/kafka"
	etcdv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
	"kratos-im/api/im"
	"kratos-im/api/social"
	"kratos-im/api/user"
	"kratos-im/app/gateway/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRedis, NewGatewayRepo, NewMQClient, NewIMServiceClient, NewRegistrar, NewDiscovery, NewSocialServiceClient, NewUserServiceClient)

// Data .
type Data struct {
	kafkaBroker  broker.Broker
	imClient     im.IMClient
	socialClient social.SocialClient
	userClient   user.UserClient
	rdb          redis.Cmdable
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rdb redis.Cmdable, kafkaBroker broker.Broker, imClient im.IMClient, socialClient social.SocialClient, userClient user.UserClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		rdb:          rdb,
		kafkaBroker:  kafkaBroker,
		imClient:     imClient,
		socialClient: socialClient,
		userClient:   userClient,
	}, cleanup, nil
}

// NewRedis 初始化redis
func NewRedis(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "shop-service/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
		DB:           int(conf.Redis.Db),
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}

	return client
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

// NewUserServiceClient user服务
func NewUserServiceClient(dis *conf.Discovery, r registry.Discovery) user.UserClient {
	conn := RpcConn(dis.Service.User, r)
	c := user.NewUserClient(conn)
	return c
}
