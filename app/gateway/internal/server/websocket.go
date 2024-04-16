package server

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"kratos-im/app/gateway/internal/conf"
	"kratos-im/app/gateway/internal/handle"
	"kratos-im/app/gateway/internal/routes"
	"kratos-im/app/gateway/internal/service"
	"kratos-im/constants"
	"kratos-im/pkg/rws"
	"net/http"
	"time"
)

// NewWebsocketServer creates a new websocket server.
func NewWebsocketServer(c *conf.Server, auth *conf.Auth, data *conf.Data, logger log.Logger, s *service.GatewayService) *rws.Server {
	token, _ := s.GenToken(auth.Key)
	svr := rws.NewServer(
		rws.WithAddr(c.Websocket.Addr),
		rws.WithMiddleware(rws.Recovery()),
		rws.WithServerDiscover(rws.NewRedisDiscover(http.Header{"Authorization": []string{fmt.Sprint("Bearer ", token)}}, constants.RedisKeyDiscoverSvr, &redis.Options{
			Addr:         data.Redis.Addr,
			Password:     data.Redis.Password,
			ReadTimeout:  data.Redis.ReadTimeout.AsDuration(),
			WriteTimeout: data.Redis.WriteTimeout.AsDuration(),
			DialTimeout:  time.Second * 2,
			PoolSize:     10,
			DB:           int(data.Redis.Db),
		})),
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
