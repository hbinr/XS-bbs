package service

import (
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/pkg/tool/snowflake"

	"github.com/gogf/gf/util/gconv"
)

// SignUp .
func (u *UserService) SignUp(param *model.SignUpParam) (dto *model.UserDto, err error) {
	var uModel dao.UserModel
	if err = u.Dao.CheckUserExist(param.Username); err != nil {
		return
	}
	if err = gconv.Struct(param, &uModel); err != nil {
		return
	}
	uModel.UserID = snowflake.GenID()
	if err = u.Dao.Insert(&uModel); err != nil {
		return
	}
	if err = gconv.Struct(uModel, &dto); err != nil {
		return
	}
	return
}

func (u *UserService) Delete(userID int64) bool {
	return u.Dao.Delete(userID)
}

func (u *UserService) Update(user *model.UserDto) error {
	var uModel model.User
	if err := gconv.Struct(user, &uModel); err != nil {
		return err
	}
	return u.Dao.Update(&uModel)
}

func (u *UserService) SlectByName(userName string) (*model.UserDto, error) {
	uModel, err := u.Dao.SlectByName(userName)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}

func (u *UserService) SelectById(userID int64) (*model.UserDto, error) {
	uModel, err := u.Dao.SelectById(userID)
	if err != nil {
		return nil, err
	}
	var uDto model.UserDto
	if err = gconv.Struct(uModel, &uDto); err != nil {
		return nil, err
	}
	return &uDto, nil
}
