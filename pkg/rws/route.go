package rws

import "github.com/gorilla/websocket"

// HandleFunc is a handle function.
type HandleFunc func(svr *Server, conn *websocket.Conn, msg Message)

// Route is a route.
type Route struct {
	// 方式
	Method string `json:"method"`
	// 处理函数
	Handle HandleFunc `json:"handle"`
}
