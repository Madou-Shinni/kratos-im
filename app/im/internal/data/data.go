package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kratos-im/app/im/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMongo, NewIMRepo, NewMQClient, NewRegistrar)

// Data .
type Data struct {
	mongoDatabase *mongo.Database
	kafkaBroker   broker.Broker
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, mongoDatabase *mongo.Database) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		mongoDatabase: mongoDatabase,
	}, cleanup, nil
}

func NewMongo(c *conf.Data, logger log.Logger) (*mongo.Database, error) {
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

	database := client.Database(c.Mongo.Db)

	return database, nil
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
