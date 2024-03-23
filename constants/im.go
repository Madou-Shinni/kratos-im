package constants

// MType 消息类型
type MType int

const (
	// MTypeText 文本消息
	MTypeText MType = iota
)

// ChatType 聊天类型
type ChatType int

const (
	// ChatTypeGroup 群聊
	ChatTypeGroup ChatType = iota + 1
	// ChatTypeSingle 单聊
	ChatTypeSingle
)
