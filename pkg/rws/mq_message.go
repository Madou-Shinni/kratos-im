package rws

import "kratos-im/constants"

type MsgChatTransfer struct {
	// 会话ID
	ConversationId string `json:"conversationId"`
	// 聊天类型
	ChatType constants.ChatType `json:"chatType"`
	// 发送者ID
	SendId string `json:"sendId"`
	// 接受者ID
	RecvId string `json:"recvId"`
	// 接受群体
	RecvIds []string `json:"recvIds"`
	// 消息体
	// 消息类型
	constants.MType `json:"msgType"`
	// 消息内容
	Content string `json:"content"`

	// 发送时间
	SendTime int64 `json:"sendTime"`
}
