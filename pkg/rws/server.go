package rws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type AckType int

const (
	AckTypeNone  AckType = iota // 无需回复
	AckTypeOnly                 // 只回复
	AckTypeRigor                // 严格回复
)

func (t AckType) ToString() string {
	switch t {
	case AckTypeNone:
		return "None"
	case AckTypeOnly:
		return "Only"
	case AckTypeRigor:
		return "Rigor"
	}

	return "None"
}

// Server is a websocket server.
type Server struct {
	*http.Server
	discover          Discover              // 发现
	routes            map[string]HandleFunc // 路由
	unaryInt          Middleware            // 中间件
	middlewares       []Middleware          // 中间件集合
	maxConnectionIdle time.Duration         // 最大连接空闲时间
	addr              string                // websocket server地址
	patten            string                // websocket connect路由
	authentication    Authentication        // 认证
	mutex             sync.RWMutex          // 读写锁
	connUserMap       map[*Conn]string      // 连接与用户的映射
	userConnMap       map[string]*Conn      // 用户与连接的映射
	upgrader          *websocket.Upgrader   // websocket升级器
	log               *log.Helper           // 日志
	ack               AckType               // 是否需要回复
	ackTimeout        time.Duration         // 回复超时时间
	sendErrCount      int                   // 发送错误次数
	concurrentCount   int                   // 并发数
	*TaskRunner                             // 任务执行器
}

// NewServer creates a new websocket server.
func NewServer(opts ...Option) *Server {
	server := &Server{
		discover:          NewNopDiscover(),
		routes:            make(map[string]HandleFunc),
		authentication:    new(DefaultAuthentication),
		connUserMap:       make(map[*Conn]string),
		userConnMap:       make(map[string]*Conn),
		maxConnectionIdle: defaultMaxConnectionIdle,
		ackTimeout:        defaultAckTimeout,
		patten:            "/ws",
		concurrentCount:   defaultConcurrentCount,
	}

	// 遍历所有的选项，并应用到Server结构体
	for _, opt := range opts {
		opt(server)
	}

	server.TaskRunner = NewTaskRunner(server.concurrentCount)

	interceptors := server.middlewares
	if len(interceptors) > 0 {
		chainUnaryServerInterceptors(server)
	}

	// 存在服务发现，采用分布式im通信的时候; 默认不做任何处理
	server.discover.Register(fmt.Sprintf("%s", server.addr))

	return server
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
func (s *Server) addConn(conn *Conn, req *http.Request) {
	uid := s.authentication.UserId(req)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 验证是否已经存在
	if c, ok := s.userConnMap[uid]; ok {
		// 关闭旧连接
		c.Close()
	}

	s.connUserMap[conn] = uid
	s.userConnMap[uid] = conn
}

func (s *Server) handleConn(conn *Conn) {

	uids := s.GetUsers(conn)
	conn.Uid = uids[0]

	// 如果存在服务发现则进行注册；默认不做任何处理
	s.discover.BoundUser(conn.Uid)

	// 处理任务
	go s.handleWrite(conn)

	// 判断是否需要ack
	if s.isAck(nil) {
		go s.readAck(conn)
	}

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

		// 根据消息进行处理
		if s.isAck(&message) {
			s.log.Infof("ack message: %v", message)
			conn.appendMsgToQueue(&message)
		} else {
			conn.messageCh <- &message
		}

	}
}

func (s *Server) isAck(message *Message) bool {
	if message == nil {
		return s.ack != AckTypeNone
	}
	return s.ack != AckTypeNone && message.FrameType != FrameAckNone && message.FrameType != FrameTranspond
}

func (s *Server) readAck(conn *Conn) {
	// 记录ack失败的次数再处理
	send := func(msg *Message, conn *Conn) error {
		err := s.SendByConns(msg, conn)
		if err == nil {
			return nil
		}

		s.log.Errorf("message ack OnlyAck send err %v message %v", err, msg)
		conn.readMessages[0].errCount++
		conn.messageMu.Unlock()

		tempDelay := time.Duration(200*conn.readMessages[0].errCount) * time.Microsecond
		if max := 1 * time.Second; tempDelay > max {
			tempDelay = max
		}

		time.Sleep(tempDelay)
		return err
	}

	for {
		select {
		case <-conn.done:
			s.log.Infof("conn done ack uid:%s", conn.Uid)
			return
		default:
		}

		// 从队列中读取消息
		conn.messageMu.Lock()
		if len(conn.readMessages) == 0 {
			conn.messageMu.Unlock()
			// 等待，让主任务更好地切换
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// 读取第一条消息
		message := conn.readMessages[0]
		if message.errCount > s.sendErrCount {
			s.log.Infof("conn send fail, message %v, ackType %v, maxSendErrCount %v", message, s.ack.ToString(), s.sendErrCount)
			conn.messageMu.Unlock()
			// 因为发送消息多次错误，而选择放弃消息
			delete(conn.readMessageSeq, message.Id)
			conn.readMessages = conn.readMessages[1:]
			continue
		}

		// 判断ack的方式
		switch s.ack {
		case AckTypeOnly:
			// 只回复
			send(&Message{FrameType: FrameAck, Id: message.Id, AckSeq: message.AckSeq + 1}, conn)
			// 业务处理
			// 从队列中删除
			conn.readMessages = conn.readMessages[1:]
			conn.messageMu.Unlock()
			conn.messageCh <- message
		case AckTypeRigor:
			// 严格回复
			// 回复
			if message.AckSeq == 0 {
				// 第一次ack
				conn.readMessages[0].AckSeq++
				conn.readMessages[0].ackTime = time.Now()
				send(&Message{FrameType: FrameAck, Id: message.Id, AckSeq: message.AckSeq}, conn)
				s.log.Infof("ack message Rigor send mid: %v, seq: %v, time: %v", message.Id, message.AckSeq, message.ackTime)
				conn.messageMu.Unlock()
				// 重发间隔
				time.Sleep(3 * time.Second)
				continue
			}

			// 验证
			// 1. 客户端返回结果，再一次确认
			msgSeq := conn.readMessageSeq[message.Id].AckSeq
			if msgSeq > message.AckSeq {
				// 确认
				conn.readMessages = conn.readMessages[1:]
				conn.messageMu.Unlock()
				conn.messageCh <- message
				s.log.Infof("ack message Rigor success send mid: %v", message.Id)
				continue
			}
			// 2. 客户端没有返回结果，是否超过超时时间
			val := s.ackTimeout - time.Since(message.ackTime)
			if !message.ackTime.IsZero() && val <= 0 {
				// 超时,删除
				s.log.Infof("ack message Rigor timeout send mid: %v", message.Id)
				delete(conn.readMessageSeq, message.Id)
				conn.readMessages = conn.readMessages[1:]
				conn.messageMu.Unlock()
				continue
			}
			// 未超时，重发
			conn.messageMu.Unlock()
			s.log.Infof("ack message Rigor resend mid: %v", message.Id)
			send(&Message{FrameType: FrameAck, Id: message.Id, AckSeq: message.AckSeq}, conn)
			// 重发间隔
			time.Sleep(3 * time.Second)
		}
	}
}

func (s *Server) dispatch(conn *Conn, message *Message, handle HandleFunc) {
	if s.unaryInt == nil {
		handle(s, conn, *message)
		return
	}
	s.unaryInt(s, conn, *message, handle)
}

func (s *Server) handleWrite(conn *Conn) {
	for {
		select {
		case <-conn.done:
			return
		case message := <-conn.messageCh:
			// 根据消息类型，调用不同的处理函数
			switch message.FrameType {
			case FramePing:
				// 心跳
				s.SendByConns(&Message{FrameType: FramePing}, conn)
			case FrameData:
				if handle, ok := s.routes[message.Method]; ok {
					s.dispatch(conn, message, handle)
				} else {
					s.log.Errorf("method not found: %s", message.Method)
				}
			case FrameTranspond:
				//s.log.Infof("server client list: %v transpond message: %v", s.connUserMap, message)
				s.SendByUsers(&Message{FrameType: FrameData, Data: message.Data}, message.TranspondUid)
			}
			if s.isAck(message) {
				// 删除队列中的消息
				conn.messageMu.Lock()
				delete(conn.readMessageSeq, message.Id)
				conn.messageMu.Unlock()
			}
		}
	}
}

func (s *Server) GetConn(uid string) *Conn {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.userConnMap[uid]
}

func (s *Server) GetConns(uids ...string) ([]*Conn, []string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var res []*Conn
	var uidsNotExist []string
	if len(uids) == 0 {
		// 获取全部
		res = make([]*Conn, 0, len(s.userConnMap))
		for _, uid := range s.userConnMap {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]*Conn, 0, len(uids))
		for _, uid := range uids {
			if _, ok := s.userConnMap[uid]; !ok {
				uidsNotExist = append(uidsNotExist, uid)
				continue
			}
			res = append(res, s.userConnMap[uid])
		}
	}

	return res, uidsNotExist
}

func (s *Server) GetUsers(conns ...*Conn) []string {

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

	conns, uids := s.GetConns(sendIds...)
	// 当前server中的连接直接发送
	s.SendByConns(msg, conns...)

	// 发现服务中转发送
	return s.discover.Transpond(msg, uids...)
}

func (s *Server) SendByConns(msg interface{}, conns ...*Conn) error {
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

func (s *Server) Close(conn *Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	uid := s.connUserMap[conn]

	if uid == "" {
		return
	}

	// 解除绑定
	s.discover.RelieveUser(uid)

	delete(s.connUserMap, conn)
	delete(s.userConnMap, uid)
	conn.Close()
}

func (s *Server) wsHandle(w http.ResponseWriter, r *http.Request) {
	// 捕获异常
	defer func() {
		if r := recover(); r != nil {
			s.log.Errorf("panic: %v", r)
		}
	}()

	conn := NewConn(s, w, r)
	if conn == nil {
		return
	}

	// 认证
	auth, err := s.authentication.Auth(w, r)
	if err != nil {
		s.log.Errorf("auth failed: %v", err)
		s.SendByConns(&Message{FrameType: FrameData, Data: fmt.Sprint("server busy")}, conn)
		conn.Close()
		return
	}

	if !auth {
		s.SendByConns(&Message{FrameType: FrameData, Data: fmt.Sprint("auth failed")}, conn)
		conn.Close()
		return
	}

	// 记录连接
	s.addConn(conn, r)
	// 处理连接
	go s.handleConn(conn)
}

// Start start the websocket server.
func (s *Server) Start(ctx context.Context) error {
	s.log.Infof("websocket server start at %s", s.addr)
	mux := http.NewServeMux()
	mux.HandleFunc(s.patten, s.wsHandle)
	s.Server = &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}
	return s.ListenAndServe()
}

// Stop stop the websocket server.
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("websocket server stop")
	return s.Shutdown(ctx)
}
