package routes

import (
	"kratos-im/app/im/internal/handle"
	"kratos-im/pkg/rws"
)

func RegisterIMWebsocketServer(srv *rws.Server) {
	srv.AddRoutes([]rws.Route{
		{Method: "user.online", Handle: handle.OnLine()},
	})
}
