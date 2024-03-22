package model

import "time"

// User 用户表
type User struct {
	Id        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	GithubId  uint64 `gorm:"column:github_id;uniqueIndex;"`        // github id
	Mobile    string `gorm:"column:mobile;uniqueIndex;size:11;"`   // 手机号
	NickName  string `gorm:"column:nickname;uniqueIndex;size:50;"` // 用户昵称
	Avatar    string `gorm:"column:avatar;size:200;"`              // 头像
}

func (u *User) TableName() string {
	return "users"
}
