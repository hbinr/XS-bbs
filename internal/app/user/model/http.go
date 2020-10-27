package model

// SignUpParam 用户注册参数
type SignUpParam struct {
	Username   string `json:"username" form:"username" v:"username@required|length:6,30#请输入用户名|用户名长度应当在:min到:max之间"`  // 用户名
	Password   string `json:"password" form:"password" v:"password@required|length:6,16#请输入密码|密码长度应当在:min到:max之间"`    // 密码
	RePassword string `json:"rePassword" form:"rePassword" v:"rePassword@required|same:password#请输入密码|两次密码不一致，请重新输入"` // 重复密码
	Nickname   string `json:"nickname" form:"nickname" v:"nickname@required#请输入中文名"`                                  // 中文名
	Email      string `json:"email" form:"email" v:"email@required|email#请输入邮箱|邮箱不合法"`                                // 邮箱
}
