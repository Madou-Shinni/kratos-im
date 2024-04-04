package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kratos-im/constants"
	"time"
)

type ChatLog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string             `bson:"conversationId"`
	SendId         string             `bson:"sendId"`
	RecvId         string             `bson:"recvId"`
	MsgFrom        int                `bson:"msgFrom"`
	MsgType        constants.MType    `bson:"msgType"`
	ChatType       constants.ChatType `bson:"chatType"`
	MsgContent     string             `bson:"msgContent"`
	SendTime       int64              `bson:"sendTime"`
	Status         int                `bson:"status"`
	ReadRecords    []byte             `bson:"readRecords"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

func (ChatLog) Collection() string {
	return "chat_log"
}
