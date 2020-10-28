package dao

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"xs.bbs/internal/pkg/constant/e"
)

var (
	ErrUserExist    = errors.New(e.CODE_USER_EXIST.Msg())
	ErrUserNotExist = errors.New(e.CODE_USER_NOT_EXIST.Msg())
	ErrEmailExist   = errors.New(e.CODE_EMAIL_EXIST.Msg())
)

// Insert 新增用户
func (u *UserDao) Insert(user *UserModel) (err error) {
	return u.DB.Create(&user).Error
}

// Delete 根据用户ID删除用户，软删除
func (u *UserDao) Delete(userID int64) bool {
	return u.DB.Where("user_id = ?", userID).Delete(&UserModel{}).RowsAffected > 0
}

// Update 根据用户ID修改用户
func (u *UserDao) Update(user *UserModel) error {
	return u.DB.Where("user_id = ?").Updates(&user).Error
}

// SelectByName 根据用户名查询用户
func (u *UserDao) SelectByName(userName string) (*UserModel, error) {
	var user UserModel
	if err := u.DB.Where("username = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// SelectByID 根据用户ID查询用户
func (u *UserDao) SelectByID(userID int64) (*UserModel, error) {
	var (
		user UserModel
		err  error
	)
	if err = u.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("userDao.SelectByID", zap.Error(err))
			return nil, ErrUserNotExist
		}
		return nil, err
	}
	return &user, nil
}

// CheckUserByUserName 根据userName检查用户是否存在
func (u *UserDao) CheckUserByUserName(userName string) error {
	var count int64
	if err := u.DB.Model(&UserModel{}).Where("username = ?", userName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExist
	}
	return nil
}

// CheckUserByEmail 通过email检查用户
func (u *UserDao) CheckUserByEmail(email string) error {
	var count int64
	if err := u.DB.Model(&UserModel{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrEmailExist
	}
	return nil
}
