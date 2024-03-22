package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"kratos-im/app/im/internal/conf"
	"net/http"
)

// WebsocketServer is a websocket server.
type WebsocketServer struct {
	addr     string
	upgrader *websocket.Upgrader
	logger   log.Logger
}

// NewWebsocketServer creates a new websocket server.
func NewWebsocketServer(c *conf.Server, logger log.Logger) *WebsocketServer {
	return &WebsocketServer{
		addr:     c.Websocket.Addr,
		upgrader: &websocket.Upgrader{},
		logger:   logger,
	}
}

// Start start the websocket server.
func (s *WebsocketServer) Start(ctx context.Context) error {
	http.HandleFunc("/ws", ConnService)
	log := log.NewHelper(log.With(s.logger, "module", "Websocket/service/websocket-service"))
	log.Infof("websocket server start at %s", s.addr)
	return http.ListenAndServe(s.addr, nil)
}

// Stop stop the websocket server.
func (s *WebsocketServer) Stop(_ context.Context) error {
	return nil
}

func ConnService(w http.ResponseWriter, r *http.Request) {
}
