package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/ginx"
)

// GetCommunityList 获取所有文章标签
func (cc *CommunityController) GetCommunityList(c *gin.Context) {
	resList, err := cc.communityService.GetCommunityList(c.Request.Context())
	if err != nil {

		ginx.RespError(c, e.CodeError)
		return
	}
	ginx.RespSuccess(c, resList)
}

// GetCommunityDetail 获取所有文章标签
func (cc *CommunityController) GetCommunityDetail(c *gin.Context) {
	var (
		id     int64
		err    error
		resDto *model.CommunityDto
	)
	if id, err = ginx.QueryInt("communityID", c); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}

	resDto, err = cc.communityService.GetCommunityDetailByID(c.Request.Context(), id)

	switch err {
	case nil:
		ginx.RespSuccess(c, resDto)
	case e.ErrInvalidID:
		ginx.RespError(c, e.CodeInvalidID)
	case e.ErrConvDataErr:
		ginx.RespError(c, e.CodeConvDataErr)
	default:
		ginx.RespError(c, e.CodeError)
	}
}
