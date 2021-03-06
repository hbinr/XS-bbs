package dao

import (
	"xs.bbs/internal/app/user/model"

	"gorm.io/gorm"
)

// 验证接口是否实现
var _ IUserDao = (*UserDao)(nil)

// UserDaoSet 使用wire 依赖注入，相当于下面的 NewUserDao 函数
//var UserDaoSet = wire.NewSet(
//	wire.Struct(new(UserDao), "*"),
//	wire.Bind(new(IUserDao), new(*UserDao)))

type (
	UserDao struct {
		db *gorm.DB
	}

	IUserDao interface {
		Insert(user *model.User) error
		Delete(int64) bool
		Update(user *model.User) error
		GetUserByID(id int64) (*model.User, error)
		GetUserByName(userName string) (*model.User, error)
		// CheckUserByUserName 根据userName检查用户是否存在
		CheckUserByUserName(userName string) error
		// CheckUserByEmail 通过用户email检查用户
		CheckUserByEmail(email string) error
	}
)

func NewUserDao(db *gorm.DB) IUserDao {
	return &UserDao{db: db}
}
