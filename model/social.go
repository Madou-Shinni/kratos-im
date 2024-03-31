package model

import (
	"database/sql"
	"kratos-im/constants"
	"time"
)

// Friends 好友关系
type Friends struct {
	ID        uint64    `gorm:"primarykey" json:"id" form:"id"` // 主键
	UserId    string    `gorm:"user_id" json:"userId"`
	FriendUid string    `gorm:"friend_uid" json:"friendUid"`
	Remark    string    `gorm:"remark" json:"remark"`
	AddSource int       `gorm:"add_source" json:"addSource"`
	CreatedAt time.Time `gorm:"created_at;type:timestamp" json:"createdAt"`
}

func (Friends) TableName() string {
	return "friends"
}

// FriendRequests 好友申请
type FriendRequests struct {
	ID           uint64                 `gorm:"primarykey" json:"id" form:"id"` // 主键
	UserId       string                 `gorm:"user_id" json:"userId"`
	ReqUid       string                 `gorm:"req_uid" json:"reqUid"`
	ReqMsg       string                 `gorm:"req_msg" json:"reqMsg"`
	ReqTime      time.Time              `gorm:"req_time;type:timestamp" json:"reqTime"`
	HandleResult constants.HandleResult `gorm:"handle_result" json:"handleResult"`
	HandleMsg    string                 `gorm:"handle_msg" json:"handleMsg"`
	HandledAt    sql.NullTime           `gorm:"handled_at" json:"handledAt"`
}

func (FriendRequests) TableName() string {
	return "friend_requests"
}

// Groups 群组
type Groups struct {
	ID              uint64    `gorm:"primarykey" json:"id" form:"id"` // 主键
	Name            string    `gorm:"name" json:"name"`
	Icon            string    `gorm:"icon" json:"icon"`
	Status          int       `gorm:"status" json:"status"`
	CreatorUid      string    `gorm:"creator_uid" json:"creatorUid"`
	GroupType       int       `gorm:"group_type" json:"groupType"`
	IsVerify        bool      `gorm:"is_verify" json:"isVerify"`
	Notification    string    `gorm:"notification" json:"notification"`
	NotificationUid string    `gorm:"notification_uid" json:"notificationUid"`
	CreatedAt       time.Time `gorm:"created_at;type:timestamp" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"updated_at;type:timestamp" json:"updatedAt"`
}

func (Groups) TableName() string {
	return "groups"
}

// GroupMembers 群成员
type GroupMembers struct {
	ID          uint64                   `gorm:"primarykey" json:"id" form:"id"` // 主键
	GroupId     uint64                   `gorm:"group_id" json:"groupId"`
	UserId      string                   `gorm:"user_id" json:"userId"`
	RoleLevel   constants.GroupRoleLevel `gorm:"role_level" json:"roleLevel"`
	JoinTime    time.Time                `gorm:"join_time;type:timestamp" json:"joinTime"`
	JoinSource  int                      `gorm:"join_source" json:"joinSource"`
	InviterUid  string                   `gorm:"inviter_uid" json:"inviterUid"`
	OperatorUid string                   `gorm:"operator_uid" json:"operatorUid"`
}

func (GroupMembers) TableName() string {
	return "group_members"
}

// GroupRequests 群申请
type GroupRequests struct {
	ID            uint64                 `gorm:"primarykey" json:"id" form:"id"` // 主键
	ReqId         string                 `gorm:"req_id" json:"reqId"`
	GroupId       uint64                 `gorm:"group_id" json:"groupId"`
	ReqMsg        string                 `gorm:"req_msg" json:"reqMsg"`
	ReqTime       sql.NullTime           `gorm:"req_time" json:"reqTime"`
	JoinSource    int                    `gorm:"join_source" json:"joinSource"`
	InviterUserId string                 `gorm:"inviter_user_id" json:"inviterUserId"`
	HandleUserId  string                 `gorm:"handle_user_id" json:"handleUserId"`
	HandleTime    sql.NullTime           `gorm:"handle_time" json:"handleTime"`
	HandleResult  constants.HandleResult `gorm:"handle_result" json:"handleResult"`
}

func (GroupRequests) TableName() string {
	return "group_requests"
}
