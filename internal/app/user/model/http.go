package model

// ParamSignUp 用户注册结构体
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
	NickName   string `json:"nickName"`
}
