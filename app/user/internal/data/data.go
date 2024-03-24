package data

import (
	"context"
	"github.com/redis/go-redis/v9"
	"kratos-im/app/user/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewRedis)

// Data .
type Data struct {
	rdb redis.Cmdable
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rdb redis.Cmdable) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		rdb: rdb,
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
