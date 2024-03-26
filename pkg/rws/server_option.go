package rws

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"time"
)

type Option func(*Server)

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

// WithMaxConnectionIdle with max connection idle time.
func WithMaxConnectionIdle(maxConnectionIdle time.Duration) Option {
	return func(s *Server) {
		s.maxConnectionIdle = maxConnectionIdle
	}
}

// WithAck with ack type.
func WithAck(ack AckType) Option {
	return func(s *Server) {
		s.ack = ack
	}
}

// WithAckTimeout with ack timeout.
func WithAckTimeout(ackTimeout time.Duration) Option {
	return func(s *Server) {
		s.ackTimeout = ackTimeout
	}
}

// WithSendErrCount with send error count.
func WithSendErrCount(sendErrCount int) Option {
	return func(s *Server) {
		s.sendErrCount = sendErrCount
	}
}
