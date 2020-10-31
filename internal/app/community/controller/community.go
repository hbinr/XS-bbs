package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/ginx"
)

// GetCommunityList 获取所有文章标签
func (a *CommunityController) GetCommunityList(c *gin.Context) {
	resList, err := a.communityService.shequ()
	if err != nil {
		ginx.ResponseError(c, e.CodeError)
		return
	}
	ginx.ResponseSuccess(c, resList)
}
