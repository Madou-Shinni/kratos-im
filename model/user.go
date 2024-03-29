package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string       `gorm:"primarykey" json:"id" form:"id"` // 主键
	CreatedAt time.Time    `json:"createdAt" form:"createdAt"`     // 创建时间
	UpdatedAt time.Time    `json:"updatedAt" form:"updatedAt"`     // 修改时间
	DeletedAt sql.NullTime `gorm:"index" json:"deletedAt" form:"deletedAt"`
	Nickname  string       `json:"nickname" gorm:"column:nickname" form:"nickname"` // 昵称
	Avatar    string       `json:"avatar" gorm:"column:avatar" form:"avatar"`       // 头像
	Email     string       `json:"email" gorm:"column:email" form:"email"`          // 邮箱
	Sex       int          `json:"sex" gorm:"column:sex" form:"sex"`                // 0:未知 1:男 2:女
	GithubId  uint         `json:"githubId" gorm:"github_id"`                       // github用户id
}
