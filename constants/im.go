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
	// ChatTypeSingle 单聊
	ChatTypeSingle = iota + 1
	// ChatTypeGroup 群聊
	ChatTypeGroup
)

// ContentType 内容类型
type ContentType int

const (
	// ContentTypeChat 聊天内容类型
	ContentTypeChat ContentType = iota
	// ContentTypeMakeRead 已读
	ContentTypeMakeRead
)
