package data

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-im/app/social/internal/biz"
	"kratos-im/app/social/internal/conf"
	"kratos-im/model"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewSocialRepo, NewTransaction)

// Data .
type Data struct {
	db *gorm.DB
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}

// NewDB 初始化db
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "social-service/data/gorm"))

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
		&model.Friends{},
		&model.FriendRequests{},
		&model.Groups{},
		&model.GroupRequests{},
		&model.GroupMembers{},
	)

	return db
}

type contextTxKey struct{}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// ExecTx gorm Transaction
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}
