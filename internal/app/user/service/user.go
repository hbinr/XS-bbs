package service

import (
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"

	"github.com/gogf/gf/util/gconv"
	"github.com/google/wire"
)

// 正常情况下会重新建一个model.UserDto结构体来做数据传输

// 验证接口是否实现
var _ IUserService = (*UserService)(nil)

type IUserService interface {
	Insert(user *model.UserDto) error
	Delete(int64) bool
	Update(user *model.UserDto) error
	SelectById(id int64) (*model.UserDto, error)
	SlectByName(userName string) (*model.UserDto, error)
}

type UserService struct {
	Dao dao.IUserDao
}

// UserServiceSet 使用 wire 依赖注入，相当于下面的 NewUserService 函数

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)))

//func NewUserService(db *gorm.DB) IUserService {
//	return &UserService{
//		Dao: dao.NewUserDao(db),
//	}
//}

func (r *UserService) Insert(user *model.UserDto) (err error) {
	var u dao.UserModel
	if err := gconv.Struct(user, &u); err != nil {
		return err
	}
	return r.Dao.Insert(&u)
}

func (r *UserService) Delete(id int64) bool {
	return r.Dao.Delete(id)
}

func (r *UserService) Update(user *model.UserDto) error {
	var uModel model.User
	if err := gconv.Struct(user, &uModel); err != nil {
		return err
	}
	return r.Dao.Update(&uModel)
}

func (r *UserService) SlectByName(userName string) (*model.UserDto, error) {
	uModel, err := r.Dao.SlectByName(userName)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}

func (r *UserService) SelectById(id int64) (*model.UserDto, error) {
	uModel, err := r.Dao.SelectById(id)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}
