package model

import (
	"gorm.io/gorm"
)

// User 用户结构体
type User struct {
	gorm.Model
	Username string `gorm:"size:32;unique;" json:"username" form:"username"` // 用户名
	Email    string `gorm:"size:128;unique;" json:"email" form:"email"`      // 邮箱
	Nickname string `gorm:"size:16;" json:"nickname" form:"nickname"`        // 昵称
	Password string `gorm:"size:512" json:"password" form:"password"`        // 密码
}
