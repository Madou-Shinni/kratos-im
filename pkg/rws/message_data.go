package rws

import "kratos-im/constants"

type (
	Msg struct {
		// 消息id
		MsgId string `mapstructure:"msgId"`
		// 消息类型
		constants.MType `mapstructure:"msgType"`
		// 消息内容
		Content string `mapstructure:"content"`
		// 已读记录
		ReadRecords map[string]string `mapstructure:"readRecords"`
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
		Msg `mapstructure:"msg"`
		// 发送时间
		SendTime int64 `mapstructure:"sendTime"`
	}

	Push struct {
		// 消息id
		MsgId string `mapstructure:"msgId"`
		// 会话ID
		ConversationId string `mapstructure:"conversationId"`
		// 聊天类型
		ChatType constants.ChatType `mapstructure:"chatType"`
		// 发送者ID
		SendId string `mapstructure:"sendId"`
		// 接受者ID
		RecvId string `mapstructure:"recvId"`
		// 接受者群体(群聊)
		RecvIds []string `mapstructure:"recvIds"`

		// 已读记录
		ReadRecords map[string]string `mapstructure:"readRecords"`
		// 内容类型
		constants.ContentType `mapstructure:"contentType"`

		// 消息体
		// 消息类型
		constants.MType `mapstructure:"msgType"`
		// 消息内容
		Content string `mapstructure:"content"`

		// 发送时间
		SendTime int64 `mapstructure:"sendTime"`
	}

	// MarkRead 已读处理
	MarkRead struct {
		// 聊天类型
		constants.ChatType `mapstructure:"chatType"`
		// 会话ID
		ConversationId string `mapstructure:"conversationId"`
		// 接受者ID
		RecvId string `mapstructure:"recvId"`
		// 消息IDs
		MsgIds []string `mapstructure:"msgIds"`
	}
)
