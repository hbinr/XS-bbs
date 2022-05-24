package service

import (
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
)

// 验证接口是否实现
var _ IUserService = (*userService)(nil)

// UserServiceSet 使用 wire 依赖注入，相当于下面的 NewUserService 函数
//var UserServiceSet = wire.NewSet(
//	wire.Struct(new(userService), "*"),
//	wire.Bind(new(IUserService), new(*userService)))

type (
	UserDto     = model.UserDto
	userService struct {
		dao dao.IUserDao
	}

	IUserService interface {
		// SignUp 注册
		SignUp(signUp *model.SignUpParam) (*UserDto, error)
		// Login 登陆
		Login(signIn *model.SignInParam) (string, error)
		Delete(int64) bool
		Update(user *model.UserDto) error
		SelectByID(id int64) (*UserDto, error)
		SelectByName(userName string) (*UserDto, error)
	}
)

func NewUserService(dao dao.IUserDao) IUserService {
	return &userService{
		dao: dao,
	}
}
