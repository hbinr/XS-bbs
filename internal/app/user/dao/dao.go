package dao

import (
	"xs.bbs/internal/app/user/model"

	"github.com/google/wire"

	"gorm.io/gorm"
)

// 验证接口是否实现
var _ IUserDao = (*UserDao)(nil)

// UserDaoSet 使用wire 依赖注入，相当于下面的 NewUserDao 函数
var UserDaoSet = wire.NewSet(
	wire.Struct(new(UserDao), "*"),
	wire.Bind(new(IUserDao), new(*UserDao)))

//func NewUserDao(db *gorm.DB) IUserDao {
//	return &UserDao{
//		DB: db,
//	}
//}

type UserModel = model.User

// IUserDao User dao 接口定义
type IUserDao interface {
	Insert(user *UserModel) error
	Delete(int64) bool
	Update(user *UserModel) error
	SelectById(id int64) (*UserModel, error)
	SlectByName(userName string) (*UserModel, error)
	CheckUserExist(userName string) error
}

type UserDao struct {
	DB *gorm.DB
}
