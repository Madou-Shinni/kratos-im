package data

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-im/app/user/internal/conf"
	"kratos-im/model"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewRedis, NewDB)

// Data .
type Data struct {
	rdb redis.Cmdable
	db  *gorm.DB
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, rdb redis.Cmdable, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		rdb: rdb,
		db:  db,
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

// NewDB 初始化db
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/gorm"))

	var config gorm.Config
	config.SkipDefaultTransaction = false
	config.DisableForeignKeyConstraintWhenMigrating = true

	dsn := conf.Database.Source

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &config)

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	db.AutoMigrate(
		&model.User{},
	)

	return db
}
