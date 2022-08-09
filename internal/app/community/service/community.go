package service

import (
	"context"

	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/pkg/constant/e"
)

// GetCommunityList 获取所有文章标签
func (s *communityService) GetCommunityList(ctx context.Context) (resList []*CommunityDto, err error) {
	var communityList []model.Community
	communityList, err = s.repo.GetCommunityList(ctx)
	for _, c := range communityList {
		dto := new(CommunityDto)
		dto.CommunityID = c.CommunityID
		dto.CommunityName = c.CommunityName
		dto.Introduction = c.Introduction

		resList = append(resList, dto)
	}
	return
}

// GetCommunityDetailByID 根据社区id获取社区详情
func (s *communityService) GetCommunityDetailByID(ctx context.Context, ID int64) (commDto *CommunityDto, err error) {
	var commuModel *model.Community

	if commuModel, err = s.repo.GetCommunityDetailByID(ctx, ID); err != nil {
		return nil, err
	}
	if err = gconv.Struct(commuModel, &commDto); err != nil {
		err = errors.Wrap(e.ErrConvDataErr, err.Error())
		return
	}
	return
}
