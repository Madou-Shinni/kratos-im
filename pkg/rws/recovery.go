package rws

import (
	"github.com/go-kratos/kratos/v2/log"
	"runtime"
)

func Recovery() Middleware {
	return func(svr *Server, conn *Conn, msg Message, handleFunc HandleFunc) {
		defer func() {
			if rerr := recover(); rerr != nil {
				buf := make([]byte, 64<<10) //nolint:gomnd
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				log.Errorf("%v: %+v\n%s\n", rerr, msg, buf)
			}
		}()
		handleFunc(svr, conn, msg)
	}
}
