package model

import (
	"gorm.io/gorm"
)

// User 用户结构体
type User struct {
	gorm.Model
	UserID   int64  `gorm:"not null;index:idx_user_id;"` // 用户ID
	Username string `gorm:"not null;size:32;unique;"`    // 用户名
	Email    string `gorm:"not null;size:128;unique;"`   // 邮箱
	Nickname string `gorm:"not null;size:16;"`           // 昵称
	Password string `gorm:"not null;size:512"`           // 密码
}
