package rws

import "time"

type FrameType uint8

const (
	FrameData      FrameType = 0x0
	FramePing      FrameType = 0x1
	FrameAck       FrameType = 0x2
	FrameAckNone   FrameType = 0x3
	FrameTranspond FrameType = 0x6 // 转发类型
	FrameErr       FrameType = 0x9
)

// Message is a message.
type Message struct {
	// 消息ID
	Id string `json:"id"`
	// 转发目标
	TranspondUid string `json:"transpondUid"`
	// 消息类型
	FrameType FrameType `json:"frameType"`
	// 方式
	Method string `json:"method"`
	// 发送者ID
	FromId string `json:"fromId"`
	// 接受者ID
	ToId string `json:"toId"`
	// 消息内容
	Data     interface{} `json:"data"`
	AckSeq   uint64      `json:"ackSeq"`
	ackTime  time.Time   `json:"ackTime"`  // 发送ack时间
	errCount int         `json:"errCount"` // 错误次数
}

// NewMessage is a constructor of Message.
func NewMessage(method, fromId, toId string, data interface{}) *Message {
	return &Message{
		FrameType: FrameData,
		Method:    method,
		FromId:    fromId,
		ToId:      toId,
		Data:      data,
	}
}

// NewErrMessage is a constructor of Err
func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
