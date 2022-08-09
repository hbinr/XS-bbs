package repository

import (
	"context"

	"xs.bbs/internal/app/user/model"

	"gorm.io/gorm"
)

// 验证接口是否实现
var _ UserRepo = (*userRepo)(nil)

// UserDaoSet 使用wire 依赖注入，相当于下面的 NewUserRepo 函数
//var UserDaoSet = wire.NewSet(
//	wire.Struct(new(userRepo), "*"),
//	wire.Bind(new(UserRepo), new(*userRepo)))

type (
	userRepo struct {
		db *gorm.DB
	}

	UserRepo interface {
		Insert(ctx context.Context, user *model.User) error
		Delete(ctx context.Context, id int64) bool
		Update(ctx context.Context, user *model.User) error
		GetUserByID(ctx context.Context, id int64) (*model.User, error)
		GetUserByName(ctx context.Context, userName string) (*model.User, error)
		// CheckUserByUserName 根据userName检查用户是否存在
		CheckUserByUserName(ctx context.Context, userName string) error
		// CheckUserByEmail 通过用户email检查用户
		CheckUserByEmail(ctx context.Context, email string) error
	}
)

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}
