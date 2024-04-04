//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-im/app/jobs/internal/biz"
	"kratos-im/app/jobs/internal/conf"
	"kratos-im/app/jobs/internal/data"
	"kratos-im/app/jobs/internal/server"
	"kratos-im/app/jobs/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.App, *conf.Registry, *conf.Discovery, *conf.MsgReadHandler, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
