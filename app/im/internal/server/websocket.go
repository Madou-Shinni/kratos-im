package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"kratos-im/app/im/internal/conf"
	"kratos-im/app/im/internal/routes"
	"kratos-im/pkg/rws"
)

// NewWebsocketServer creates a new websocket server.
func NewWebsocketServer(c *conf.Server, logger log.Logger) *rws.Server {
	svr := rws.NewServer(
		rws.WithAddr(c.Websocket.Addr),
		rws.WithPatten("/ws"),
		rws.WithLogger(log.NewHelper(log.With(logger, "module", "Websocket/service/websocket-service"))),
		rws.WithUpgrader(&websocket.Upgrader{}),
	)

	routes.RegisterIMWebsocketServer(svr)
	return svr
}
