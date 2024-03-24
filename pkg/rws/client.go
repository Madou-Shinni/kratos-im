package rws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
)

type IClient interface {
	// Close 关闭连接
	Close() error
	// Send 发送消息
	Send(v any) error
	// Read 接收消息
	Read(v any) error
}

type client struct {
	*websocket.Conn
	host   string
	patten string
	header http.Header
}

func NewClient(host, patten string, header http.Header) (IClient, error) {
	c := &client{
		Conn:   nil,
		host:   host,
		patten: patten,
		header: header,
	}

	conn, err := c.dial()
	if err != nil {
		return nil, err
	}

	c.Conn = conn

	return c, nil
}

func (c *client) dial() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: c.host, Path: c.patten}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.header)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = c.Conn.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}

	// 有错误 再增加一个重连发送
	dial, err := c.dial()
	if err != nil {
		return err
	}

	c.Conn = dial

	return c.Conn.WriteMessage(websocket.TextMessage, data)
}

func (c *client) Read(v any) error {
	_, msg, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}

	return json.Unmarshal(msg, &v)
}
