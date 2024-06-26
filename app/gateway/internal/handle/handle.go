package handle

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"kratos-im/app/gateway/internal/service"
	"kratos-im/constants"
	"kratos-im/pkg/rws"
	"kratos-im/pkg/tools"
	"time"
)

// OnLine 上线
func OnLine(s *service.GatewayService) rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		uids := svr.GetUsers()
		myids := svr.GetUsers(conn)
		if len(uids) == 0 {
			return
		}
		// 发送给所有人
		me := myids[0]
		for _, uid := range uids {
			if uid != me && uid != constants.RedisKeyOnlineUser {
				svr.SendByUsers(rws.NewMessage("online", me, uid, me+" 上线啦!"), uids...)
			}
		}
	}
}

// Chat 聊天
func Chat(s *service.GatewayService) rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		var data rws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}

		if data.ConversationId == "" {
			switch data.ChatType {
			case constants.ChatTypeSingle:
				data.ConversationId = tools.CombineId(conn.Uid, data.RecvId)
			case constants.ChatTypeGroup:
				data.ConversationId = data.RecvId
			}
		}

		// 基于kafka的异步消息处理
		err := s.KafkaBroker.Publish(context.Background(), constants.TopicMsgTransfer, rws.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MType:          data.MType,
			Content:        data.Content,
			SendTime:       time.Now().UnixMilli(),
		})
		if err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}
	}
}

// MakeRead 消息已读
func MakeRead(s *service.GatewayService) rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		var data rws.MarkRead
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}

		// 基于kafka的异步消息处理
		err := s.KafkaBroker.Publish(context.Background(), constants.TopicMsgReadTransfer, rws.MsgMarkReadTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MsgIds:         data.MsgIds,
		})
		if err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}
	}
}

// Push 消息推送
func Push(s *service.GatewayService) rws.HandleFunc {
	return func(svr *rws.Server, conn *rws.Conn, msg rws.Message) {
		var data rws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			svr.SendByConns(rws.NewErrMessage(err), conn)
			return
		}
		// 发送消息
		switch data.ChatType {
		case constants.ChatTypeSingle: // 私聊
			pushSingle(svr, data, data.RecvId)
		case constants.ChatTypeGroup: // 群聊
			pushGroup(svr, data)
		}
	}
}

// 私聊推送
func pushSingle(svr *rws.Server, data rws.Push, recvId string) error {
	conn := svr.GetConn(recvId)
	if conn == nil {
		// 对方不在线
		return svr.SendByUsers(rws.NewMessage("push", data.SendId, data.RecvId, rws.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendTime:       data.SendTime,
			Msg: rws.Msg{
				MsgId:       data.MsgId,
				MType:       data.MType,
				Content:     data.Content,
				ReadRecords: data.ReadRecords,
			},
		}), recvId)
	}
	// 发送消息
	return svr.SendByConns(rws.NewMessage("push", data.SendId, data.RecvId, rws.Chat{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		SendTime:       data.SendTime,
		Msg: rws.Msg{
			MsgId:       data.MsgId,
			MType:       data.MType,
			Content:     data.Content,
			ReadRecords: data.ReadRecords,
		},
	}), conn)
}

// 群聊推送
func pushGroup(svr *rws.Server, data rws.Push) error {
	// 并发发送
	for _, recvId := range data.RecvIds {
		func(id string) {
			svr.Schedule(func() {
				pushSingle(svr, data, id)
			})
		}(recvId)
	}
	return nil
}
