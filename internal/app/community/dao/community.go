package dao

import (
	"go.uber.org/zap"
	"xs.bbs/internal/app/community/model"
)

// GetCommunityList 获取所有文章标签
func (d *CommunityDao) GetCommunityList() ([]model.Community, error) {
	var resList []model.Community
	if err := d.DB.Find(&resList).Error; err != nil {
		zap.L().Error("Dao.GetCommunityList failed", zap.Error(err))
		return nil, err
	}
	return resList, nil
}
