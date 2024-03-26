package rws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type Conn struct {
	idleMu sync.Mutex
	Uid    string
	mu     sync.Mutex
	*websocket.Conn
	s                 *Server
	idle              time.Time     // 上次空闲时间
	maxConnectionIdle time.Duration // 最大连接空闲时间
	messageMu         sync.Mutex
	readMessages      []*Message
	readMessageSeq    map[string]*Message
	messageCh         chan *Message

	done chan struct{} // 连接关闭信号
}

func NewConn(s *Server, w http.ResponseWriter, r *http.Request) *Conn {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.Errorf("upgrade: %v", err)
		return nil
	}

	conn := &Conn{
		Conn:              c,
		s:                 s,
		idle:              time.Now(),
		maxConnectionIdle: s.maxConnectionIdle,
		readMessageSeq:    make(map[string]*Message, 2),
		readMessages:      make([]*Message, 0, 2),
		messageCh:         make(chan *Message, 1),
		done:              make(chan struct{}),
	}

	go conn.keepalive()

	return conn
}

func (c *Conn) appendMsgToQueue(msg *Message) {
	c.messageMu.Lock()
	defer c.messageMu.Unlock()

	// 读队列中
	if m, ok := c.readMessageSeq[msg.Id]; ok {
		// 消息已存在，已经有ack确认
		if len(c.readMessages) == 0 {
			// 队列中没有消息
			return
		}

		if m.AckSeq >= msg.AckSeq {
			// 没有进行ack确认，或者重复
			return
		}

		c.readMessageSeq[msg.Id] = msg

		return
	}

	// 没有ack确认
	if msg.FrameType == FrameAck {
		// 避免ack消息重复
		return
	}

	c.readMessages = append(c.readMessages, msg)
	c.readMessageSeq[msg.Id] = msg
}

func (c *Conn) ReadMessage() (messageType int, p []byte, err error) {
	messageType, p, err = c.Conn.ReadMessage()

	// 读取完成 进入工作时间
	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	c.idle = time.Time{}
	return
}

func (c *Conn) WriteMessage(messageType int, data []byte) error {
	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	err := c.Conn.WriteMessage(messageType, data)
	// 写入完成 进入空闲时间
	c.idle = time.Now()
	return err
}

func (c *Conn) Close() error {
	select {
	case <-c.done: // 连接已关闭 不再关闭
	default:
		close(c.done)
	}

	return c.Conn.Close()
}

// 长连接检测机制
func (c *Conn) keepalive() {
	idleTimer := time.NewTimer(c.maxConnectionIdle)
	defer idleTimer.Stop()

	for {
		select {
		case <-idleTimer.C:
			c.idleMu.Lock()
			idle := c.idle

			//fmt.Printf("idle %v, maxIdle %v \n", c.idle, c.maxConnectionIdle)
			if idle.IsZero() { // The connection is non-idle.
				c.idleMu.Unlock()
				idleTimer.Reset(c.maxConnectionIdle)
				continue
			}
			val := c.maxConnectionIdle - time.Since(idle)
			//fmt.Printf("val %v \n", val)
			c.idleMu.Unlock()
			if val <= 0 {
				// The connection has been idle for a duration of keepalive.MaxConnectionIdle or more.
				// Gracefully close the connection.
				c.s.Close(c)
				return
			}
			idleTimer.Reset(val)
		case <-c.done:
			fmt.Println("客户端结束连接")
			return
		}
	}
}
