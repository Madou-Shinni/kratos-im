package handle

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"kratos-im/app/im/internal/service"
	"kratos-im/constants"
	"kratos-im/pkg/rws"
)

// OnLine 上线
func OnLine(s *service.IMService) rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		uids := svr.GetUsers()
		myids := svr.GetUsers(conn)
		if len(uids) == 0 {
			return
		}
		// 发送给所有人
		me := myids[0]
		for _, uid := range uids {
			if uid != me && uid != constants.SystemRootUid {
				svr.SendByUsers(rws.NewMessage("online", me, uid, me+" 上线啦!"), uids...)
			}
		}
	}
}

// Chat 私聊
func Chat(s *service.IMService) rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		var data rws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}
		//// 保存聊天记录
		//err := s.CreateChatLog(context.Background(), &data, conn.Uid)
		//if err != nil {
		//	svr.SendByConns(rws.NewErrMessage(err), conn)
		//	return
		//}
		//// 发送给对方
		//svr.SendByUsers(rws.NewMessage("chat", conn.Uid, data.RecvId, rws.Chat{
		//	ConversationId: data.ConversationId,
		//	ChatType:       data.ChatType,
		//	SendId:         conn.Uid,
		//	RecvId:         data.RecvId,
		//	Msg:            data.Msg,
		//	SendTime:       time.Now().UnixMilli(),
		//}), data.RecvId)

		// 基于kafka的异步消息处理
		err := s.KafkaBroker.Publish(context.Background(), constants.TopicMsgTransfer, rws.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MType:          data.MType,
			Content:        data.Content,
			SendTime:       0,
		})
		if err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}
	}
}

// Push 消息推送
func Push(s *service.IMService) rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		var data rws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}
		// 发送消息
		recvid := svr.GetConn(data.RecvId)
		if recvid == nil {
			// TODO: 对方不在线

			return
		}

		// 发送给对方
		svr.SendByConns(rws.NewMessage("push", conn.Uid, data.RecvId, rws.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendTime:       0,
			Msg: rws.Msg{
				MType:   data.MType,
				Content: data.Content,
			},
		}), recvid)
	}
}
