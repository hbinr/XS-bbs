package model

// User 用户结构体
type UserDto struct {
	Username string `json:"username" form:"username"` // 用户名
	Email    string `json:"email" form:"email"`       // 邮箱
	Nickname string `json:"nickname" form:"nickname"` // 昵称
	Password string `json:"password" form:"password"` // 密码
}
