package handle

import (
	"kratos-im/pkg/rws"
)

// OnLine 上线
func OnLine() rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		uids := svr.GetUsers()
		myids := svr.GetUsers(conn)
		if len(uids) == 0 {
			return
		}
		// 发送给所有人
		me := myids[0]
		for _, uid := range uids {
			svr.SendByConns(rws.NewMessage("online", me, uid, me+" 上线啦!"), conn)
		}
	}
}
