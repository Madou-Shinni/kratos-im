package routes

import (
	"kratos-im/app/im/internal/handle"
	"kratos-im/app/im/internal/service"
	"kratos-im/pkg/rws"
)

func RegisterIMWebsocketServer(srv *rws.Server, s *service.IMService) {
	srv.AddRoutes([]rws.Route{
		{Method: "user.online", Handle: handle.OnLine(s)},
		{Method: "user.chat", Handle: handle.Chat(s)},
	})
}
