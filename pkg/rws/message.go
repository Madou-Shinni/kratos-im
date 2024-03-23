package rws

type FrameType uint8

const (
	FrameData FrameType = 0x0
	FramePing FrameType = 0x1
)

// Message is a message.
type Message struct {
	// 消息类型
	FrameType FrameType `json:"frameType"`
	// 方式
	Method string `json:"method"`
	// 发送者ID
	FromId string `json:"fromId"`
	// 接受者ID
	ToId string `json:"toId"`
	// 消息内容
	Data interface{} `json:"data"`
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
