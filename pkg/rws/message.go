package rws

// Message is a message.
type Message struct {
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
		Method: method,
		FromId: fromId,
		ToId:   toId,
		Data:   data,
	}
}
