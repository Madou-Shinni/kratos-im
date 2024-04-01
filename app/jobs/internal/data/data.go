package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kratos-im/app/jobs/internal/conf"
	"kratos-im/constants"
	"kratos-im/pkg/rws"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewConsumerRepo, NewMongo, NewRedis, NewWsClient, NewRegistrar)

// Data .
type Data struct {
	rdb           redis.Cmdable
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
	wsClient      rws.IClient
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rdb redis.Cmdable, mongoClient *mongo.Client, wsClient rws.IClient) (*Data, func(), error) {
	data := &Data{
		rdb:           rdb,
		mongoDatabase: mongoClient.Database(c.Mongo.Db),
		wsClient:      wsClient,
	}
	cleanup := func() {
		wsClient.Close()
		mongoClient.Disconnect(context.Background())
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
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

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
