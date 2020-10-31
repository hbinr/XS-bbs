package service

import (
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
	"xs.bbs/internal/app/community/model"
)

// GetCommunityList 获取所有文章标签
func (s *CommunityService) GetCommunityList() (resList []model.CommunityDto, err error) {
	var communityList []model.Community
	communityList, err = s.Dao.GetCommunityList()
	for _, c := range communityList {
		var dto model.CommunityDto
		if err = gconv.Struct(c, &dto); err != nil {
			zap.L().Error("CommunityService.GetCommunityList->gconv.Struct failed", zap.Error(err))
			return nil, err
		}
		resList = append(resList, dto)
	}
	return resList, nil
}
