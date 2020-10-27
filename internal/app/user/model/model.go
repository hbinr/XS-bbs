package model

import (
	"gorm.io/gorm"
)

// User 用户结构体
type User struct {
	gorm.Model
	UserID   int64  `gorm:"not null;index:idx_user_id;"` // 用户ID
	Username  string `gorm:"size:32;unique;"`             // 用户名
	Email    string `gorm:"size:128;unique;"`            // 邮箱
	Nickname string `gorm:"size:16;"`                    // 昵称
	Password string `gorm:"size:512"`                    // 密码
}
