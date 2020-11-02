package dao

import (
	"errors"

	"xs.bbs/internal/pkg/constant/e"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (c *CommunityDao) GetCommunityList() (resList []CommunityModel, err error) {
	if err = c.db.Find(&resList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("dao.GetCommunityList no data", zap.Error(err))
			err = gorm.ErrRecordNotFound
		}
		zap.L().Error("dao.GetCommunityList failed", zap.Error(err))
	}
	return resList, err
}

func (c *CommunityDao) GetCommunityDetailByID(ID int64) (res *CommunityModel, err error) {
	res = new(CommunityModel)
	if err = c.db.Where("community_id", ID).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("dao.GetCommunityList no data", zap.Error(err))
			err = e.ErrInvalidID
		}
		zap.L().Error("dao.GetCommunityList failed", zap.Error(err))
	}
	return res, err
}
