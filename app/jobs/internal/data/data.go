package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	ggrpc "google.golang.org/grpc"
	"kratos-im/api/social"
	"kratos-im/app/jobs/internal/conf"
	"kratos-im/constants"
	"kratos-im/pkg/rws"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewConsumerRepo, NewMongo, NewRedis, NewWsClient, NewRegistrar, NewDiscovery, NewSocialServiceClient)

// Data .
type Data struct {
	rdb           redis.Cmdable
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
	wsClient      rws.IClient
	socialClient  social.SocialClient
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rdb redis.Cmdable, mongoClient *mongo.Client, wsClient rws.IClient, socialClient social.SocialClient) (*Data, func(), error) {
	data := &Data{
		rdb:           rdb,
		mongoDatabase: mongoClient.Database(c.Mongo.Db),
		wsClient:      wsClient,
		socialClient:  socialClient,
	}
	cleanup := func() {
		wsClient.Close()
		mongoClient.Disconnect(context.Background())
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
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

func NewMongo(c *conf.Data, logger log.Logger) (*mongo.Client, error) {
	log := log.NewHelper(logger)
	//1.建立连接
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.Mongo.Url).SetConnectTimeout(5*time.Second))
	if err != nil {
		// 连接失败
		log.Error(err)
		return nil, err
	}

	// Ping the MongoDB to ensure connectivity
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error("Failed to ping MongoDB:", err)
		return nil, err
	}

	return client, nil
}

// NewRedis 初始化redis
func NewRedis(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "jobs-service/data/redis"))
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

func NewWsClient(c *conf.Data, logger log.Logger, rdb redis.Cmdable) (rws.IClient, error) {
	log := log.NewHelper(logger)
	//1.建立连接
	token, err := rdb.Get(context.Background(), constants.SystemRootUid).Result()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	header := http.Header{}
	header.Add("Authorization", fmt.Sprint("Bearer ", token))
	client, err := rws.NewClient(c.Ws.Host, c.Ws.Patten, header)
	if err != nil {
		// 连接失败
		log.Error(err)
		return nil, err
	}

	return client, nil
}

// NewSocialServiceClient social服务
func NewSocialServiceClient(dis *conf.Discovery, r registry.Discovery) social.SocialClient {
	conn := RpcConn(dis.Service.Social, r)
	c := social.NewSocialClient(conn)
	return c
}
