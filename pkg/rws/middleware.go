package rws

type Middleware func(svr *Server, conn *Conn, msg Message, handleFunc HandleFunc)

func chainUnaryServerInterceptors(svr *Server) {
	var chainedInt Middleware
	interceptors := svr.middlewares
	if len(interceptors) == 0 {
		chainedInt = nil
	} else if len(interceptors) == 1 {
		chainedInt = interceptors[0]
	} else {
		chainedInt = chainUnaryInterceptors(interceptors)
	}
	svr.unaryInt = chainedInt
}

func chainUnaryInterceptors(interceptors []Middleware) Middleware {
	return func(svr *Server, conn *Conn, msg Message, handleFunc HandleFunc) {
		interceptors[0](svr, conn, msg, getChainUnaryHandler(interceptors, 1, handleFunc))
	}
}

func getChainUnaryHandler(interceptors []Middleware, curr int, finalHandler HandleFunc) HandleFunc {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(svr *Server, conn *Conn, msg Message) {
		interceptors[curr+1](svr, conn, msg, getChainUnaryHandler(interceptors, curr+1, finalHandler))
	}
}
