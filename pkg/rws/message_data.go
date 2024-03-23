package rws

import "kratos-im/constants"

type (
	Msg struct {
		// 消息类型
		constants.MType `mapstructure:"msgType"`
		// 消息内容
		Content string `mapstructure:"content"`
	}

	Chat struct {
		// 会话ID
		ConversationId string `mapstructure:"conversationId"`
		// 聊天类型
		ChatType constants.ChatType `mapstructure:"chatType"`
		// 发送者ID
		SendId string `mapstructure:"sendId"`
		// 接受者ID
		RecvId string `mapstructure:"recvId"`
		// 消息体
		Msg Msg `mapstructure:"msg"`
		// 发送时间
		SendTime int64 `mapstructure:"sendTime"`
	}
)
