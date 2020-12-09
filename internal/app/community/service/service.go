package service

import (
	"xs.bbs/internal/app/community/dao"
	"xs.bbs/internal/app/community/model"
)

var _ ICommunityService = (*communityService)(nil)

// CommunityServiceSet CommunityServiceSet依赖注入
//var CommunityServiceSet = wire.NewSet(
//	wire.Struct(new(CommunityService), "*"),
//	wire.Bind(new(ICommunityService), new(*CommunityService)),
//)

type (
	CommunityDto     = model.CommunityDto
	communityService struct {
		dao dao.ICommunityDao
	}

	// ICommunityService 文章标签接口
	ICommunityService interface {
		// GetCommunityList 获取所有文章标签
		GetCommunityList() ([]CommunityDto, error)
		// GetCommunityDetailByID 根据社区id获取社区详情
		GetCommunityDetailByID(int64) (*CommunityDto, error)
	}
)

func NewCommunityService(dao dao.ICommunityDao) ICommunityService {
	return &communityService{
		dao: dao,
	}
}
