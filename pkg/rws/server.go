package rws

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Option func(*Server)

// Server is a websocket server.
type Server struct {
	routes         map[string]HandleFunc      // 路由
	addr           string                     // websocket server地址
	patten         string                     // websocket connect路由
	authentication Authentication             // 认证
	mutex          sync.RWMutex               // 读写锁
	connUserMap    map[*websocket.Conn]string // 连接与用户的映射
	userConnMap    map[string]*websocket.Conn // 用户与连接的映射
	upgrader       *websocket.Upgrader        // websocket升级器
	log            *log.Helper                // 日志
}

// NewServer creates a new websocket server.
func NewServer(opts ...Option) *Server {
	server := &Server{
		routes:         make(map[string]HandleFunc),
		authentication: new(DefaultAuthentication),
		connUserMap:    make(map[*websocket.Conn]string),
		userConnMap:    make(map[string]*websocket.Conn),
	}

	// 遍历所有的选项，并应用到Server结构体
	for _, opt := range opts {
		opt(server)
	}

	return server
}

// WithAddr with server address.
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// WithUpgrader with websocket upgrader.
func WithUpgrader(upgrader *websocket.Upgrader) Option {
	return func(s *Server) {
		s.upgrader = upgrader
	}
}

// WithLogger with logger.
func WithLogger(logger *log.Helper) Option {
	return func(s *Server) {
		s.log = logger
	}
}

// WithAuthentication with authentication.
func WithAuthentication(authentication Authentication) Option {
	return func(s *Server) {
		s.authentication = authentication
	}
}

// WithPatten with patten.
func WithPatten(patten string) Option {
	return func(s *Server) {
		s.patten = patten
	}
}

// AddRoutes add routes.
func (s *Server) AddRoutes(routes []Route) {
	for _, route := range routes {
		if _, ok := s.routes[route.Method]; ok {
			panic("Route method already exists: " + route.Method)
		}
		s.routes[route.Method] = route.Handle
	}
}

// addConn .
func (s *Server) addConn(conn *websocket.Conn, req *http.Request) {
	uid := s.authentication.UserId(req)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.connUserMap[conn] = uid
	s.userConnMap[uid] = conn
}

func (s *Server) handleConn(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.log.Errorf("read message err: %v", err)
			s.Close(conn)
			return
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			s.log.Errorf("unmarshal message err: %v", err)
			s.Close(conn)
			return
		}

		// 根据消息类型，调用不同的处理函数
		if handle, ok := s.routes[message.Method]; ok {
			handle(s, conn, message)
		} else {
			s.log.Errorf("method not found: %s", message.Method)
		}

	}
}

func (s *Server) GetConn(uid string) *websocket.Conn {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.userConnMap[uid]
}

func (s *Server) GetConns(uids ...string) []*websocket.Conn {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var res []*websocket.Conn
	if len(uids) == 0 {
		// 获取全部
		res = make([]*websocket.Conn, 0, len(s.userConnMap))
		for _, uid := range s.userConnMap {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]*websocket.Conn, 0, len(uids))
		for _, uid := range uids {
			res = append(res, s.userConnMap[uid])
		}
	}

	return res
}

func (s *Server) GetUsers(conns ...*websocket.Conn) []string {

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connUserMap))
		for _, uid := range s.connUserMap {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connUserMap[conn])
		}
	}

	return res
}

func (s *Server) SendByUsers(msg interface{}, sendIds ...string) error {
	if len(sendIds) == 0 {
		return nil
	}

	return s.SendByConns(msg, s.GetConns(sendIds...)...)
}

func (s *Server) SendByConns(msg interface{}, conns ...*websocket.Conn) error {
	if len(conns) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) Close(conn *websocket.Conn) {
	conn.Close()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	uid := s.connUserMap[conn]
	delete(s.connUserMap, conn)
	delete(s.userConnMap, uid)
}

func (s *Server) wsHandle(w http.ResponseWriter, r *http.Request) {
	// 捕获异常
	defer func() {
		if r := recover(); r != nil {
			s.log.Errorf("panic: %v", r)
		}
	}()

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.Errorf("upgrade: %v", err)
		return
	}

	// 认证
	auth, err := s.authentication.Auth(w, r)
	if err != nil {
		s.log.Errorf("auth failed: %v", err)
		conn.WriteMessage(websocket.CloseInternalServerErr, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "auth failed"))
		return
	}

	if !auth {
		s.log.Error("auth failed")
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "auth failed"))
		return
	}

	// 记录连接
	s.addConn(conn, r)
	// 处理连接
	go s.handleConn(conn)
}

// Start start the websocket server.
func (s *Server) Start(ctx context.Context) error {
	http.HandleFunc(s.patten, s.wsHandle)
	s.log.Infof("websocket server start at %s", s.addr)
	return http.ListenAndServe(s.addr, nil)
}

// Stop stop the websocket server.
func (s *Server) Stop(_ context.Context) error {
	return nil
}
