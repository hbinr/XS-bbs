package dao

import (
	"errors"

	"gorm.io/gorm"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/pkg/log"
)

func (r *UserDao) Insert(user *UserModel) (err error) {
	return r.DB.Create(&user).Error
}

func (r *UserDao) Delete(userID int64) bool {
	return r.DB.Where("user_id = ?", userID).Delete(&UserModel{}).RowsAffected > 0
}

func (r *UserDao) Update(user *UserModel) error {
	return r.DB.Where("user_id = ?").Updates(&user).Error
}

func (r *UserDao) SlectByName(userName string) (*UserModel, error) {
	var user UserModel
	if err := r.DB.Where("username = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserDao) SelectById(userID int64) (*UserModel, error) {
	var (
		user UserModel
		err  error
	)
	if err = r.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("userDao.SelectById", err)
			return nil, errors.New(e.GetMsg(e.ERROR_NOT_EXIST_USER))
		}
		return nil, err
	}
	return &user, nil
}

// CheckUserExist 根据userName检查用户是否存在
func (r *UserDao) CheckUserExist(userName string) error {
	var count int64
	if err := r.DB.Model(&UserModel{}).Where("username = ?", userName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New(e.GetMsg(e.ERROR_EXIST_USER))
	}
	return nil
}
