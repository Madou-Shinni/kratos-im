package routes

import (
	"kratos-im/app/gateway/internal/handle"
	"kratos-im/app/gateway/internal/service"
	"kratos-im/pkg/rws"
)

func RegisterIMWebsocketServer(srv *rws.Server, s *service.GatewayService) {
	srv.AddRoutes([]rws.Route{
		{Method: "user.online", Handle: handle.OnLine(s)},
		{Method: "user.chat", Handle: handle.Chat(s)},
		{Method: "push", Handle: handle.Push(s)},
	})
}
