package rws

import "net/http"

type ClientOption func(o *client)

// WithHeader 设置请求头
func WithHeader(header http.Header) ClientOption {
	return func(c *client) {
		c.header = header
	}
}

// WithHost 设置主机地址
func WithHost(host string) ClientOption {
	return func(c *client) {
		c.host = host
	}
}

// WithClientPatten 设置路径
func WithClientPatten(patten string) ClientOption {
	return func(c *client) {
		c.patten = patten
	}
}
