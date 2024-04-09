package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"kratos-im/app/gateway/internal/conf"
	"kratos-im/app/gateway/internal/handle"
	"kratos-im/app/gateway/internal/routes"
	"kratos-im/app/gateway/internal/service"
	"kratos-im/pkg/rws"
	"time"
)

// NewWebsocketServer creates a new websocket server.
func NewWebsocketServer(c *conf.Server, auth *conf.Auth, logger log.Logger, s *service.GatewayService) *rws.Server {
	svr := rws.NewServer(
		rws.WithAddr(c.Websocket.Addr),
		rws.WithMiddleware(rws.Recovery()),
		rws.WithPatten("/ws"),
		rws.WithLogger(log.NewHelper(log.With(logger, "module", "Websocket/service/websocket-service"))),
		rws.WithUpgrader(&websocket.Upgrader{}),
		rws.WithSendErrCount(3),
		rws.WithAck(rws.AckTypeNone),
		rws.WithAckTimeout(10*time.Second),
		//rws.WithMaxConnectionIdle(time.Second*10),
		rws.WithAuthentication(handle.NewJWTAuth(func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.Key), nil
		}, handle.WithClaims(func() jwt.Claims {
			return jwt.MapClaims{}
		}))),
	)

	routes.RegisterIMWebsocketServer(svr, s)
	return svr
}
