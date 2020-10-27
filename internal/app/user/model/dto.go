package model

// UserDto .
type UserDto struct {
	ID       uint   `json:"id"`                       // ID
	UserID   int64  `json:"userID"`                   // 用户ID
	Username string `json:"username" form:"username"` // 用户名
	Email    string `json:"email" form:"email"`       // 邮箱
	Nickname string `json:"nickname" form:"nickname"` // 昵称
}
