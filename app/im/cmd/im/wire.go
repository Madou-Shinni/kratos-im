//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-im/app/im/internal/biz"
	"kratos-im/app/im/internal/conf"
	"kratos-im/app/im/internal/data"
	"kratos-im/app/im/internal/server"
	"kratos-im/app/im/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Auth, *conf.App, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
	//panic(wire.Build(server.ProviderSet, newApp))
}
