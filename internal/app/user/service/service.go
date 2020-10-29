package service

import (
	"github.com/google/wire"
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
)

// 验证接口是否实现
var _ IUserService = (*UserService)(nil)

// UserServiceSet 使用 wire 依赖注入，相当于下面的 NewUserService 函数
var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)))

type UserService struct {
	Dao dao.IUserDao
}
type IUserService interface {
	// SignUp 注册
	SignUp(signUp *model.SignUpParam) (*model.UserDto, error)
	// SignIn 登陆
	SignIn(signIn *model.SignInParam) (string, error)
	Delete(int64) bool
	Update(user *model.UserDto) error
	SelectByID(id int64) (*model.UserDto, error)
	SelectByName(userName string) (*model.UserDto, error)
}
