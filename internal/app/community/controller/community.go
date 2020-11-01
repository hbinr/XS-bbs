package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/ginx"
)

// GetCommunityList 获取所有文章标签
func (a *CommunityController) GetCommunityList(c *gin.Context) {
	resList, err := a.communityService.GetCommunityList()
	if err != nil {

		ginx.ResponseError(c, e.CodeError)
		return
	}
	ginx.ResponseSuccess(c, resList)
}

// GetCommunityList 获取所有文章标签
func (a *CommunityController) GetCommunityDetail(c *gin.Context) {
	var (
		id     int64
		err    error
		resDto *model.CommunityDto
	)
	if id, err = ginx.QueryInt("communityID", c); err != nil {
		ginx.ResponseError(c, e.CodeInvalidParams)
		return
	}

	resDto, err = a.communityService.GetCommunityDetailByID(id)
	if err != nil {
		if errors.Is(err, e.ErrInvalidID) {
			ginx.ResponseError(c, e.CodeInvalidID)
			return
		}
		ginx.ResponseError(c, e.CodeError)
		return
	}
	ginx.ResponseSuccess(c, resDto)
}
