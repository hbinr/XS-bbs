package service

import (
	"context"

	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/app/user/repository"
)

// 验证接口是否实现
var _ UserService = (*userService)(nil)

// UserServiceSet 使用 wire 依赖注入，相当于下面的 NewUserService 函数
//var UserServiceSet = wire.NewSet(
//	wire.Struct(new(userService), "*"),
//	wire.Bind(new(UserService), new(*userService)))

type (
	UserDto     = model.UserDto
	userService struct {
		repo repository.UserRepo
	}

	UserService interface {
		// Register 注册
		Register(ctx context.Context, signUp *model.RegisterReq) (*UserDto, error)
		// Login 登陆
		Login(ctx context.Context, signIn *model.LoginReq) (string, error)
		Delete(ctx context.Context, id int64) bool
		Update(ctx context.Context, user *model.UserDto) error
		SelectByID(ctx context.Context, id int64) (*UserDto, error)
		SelectByName(ctx context.Context, userName string) (*UserDto, error)
	}
)

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}
