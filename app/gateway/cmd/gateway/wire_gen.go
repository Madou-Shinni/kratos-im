// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-im/app/gateway/internal/biz"
	"kratos-im/app/gateway/internal/conf"
	"kratos-im/app/gateway/internal/data"
	"kratos-im/app/gateway/internal/server"
	"kratos-im/app/gateway/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, discovery *conf.Discovery, app *conf.App, registry *conf.Registry, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	broker := data.NewMQClient(confData, logger)
	registryDiscovery := data.NewDiscovery(registry)
	imClient := data.NewIMServiceClient(discovery, registryDiscovery)
	socialClient := data.NewSocialServiceClient(discovery, registryDiscovery)
	dataData, cleanup, err := data.NewData(confData, logger, broker, imClient, socialClient)
	if err != nil {
		return nil, nil, err
	}
	gatewayRepo := data.NewGatewayRepo(dataData, logger)
	gatewayUsecase := biz.NewGatewayUsecase(gatewayRepo, logger)
	gatewayService := service.NewGatewayService(gatewayUsecase)
	httpServer := server.NewHTTPServer(confServer, auth, gatewayService, logger)
	rwsServer := server.NewWebsocketServer(confServer, auth, logger, gatewayService)
	registrar := data.NewRegistrar(registry)
	kratosApp := newApp(logger, httpServer, rwsServer, app, registrar)
	return kratosApp, func() {
		cleanup()
	}, nil
}
