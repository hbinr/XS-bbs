package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/post/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/ginx"
)

func (p *PostController) Create(c *gin.Context) {
	var (
		err       error
		userID    int64
		postParam model.PostParam
		postDto   *model.PostDto
	)
	if errStr := ginx.BindAndValid(c, &postParam); errStr != "" {
		ginx.ResponseErrorWithMsg(c, e.CodeError, errStr)
		return
	}
	if userID, err = ginx.GetCurrentUserID(c); err != nil {
		ginx.ResponseError(c, e.CodeNeedLogin)
		return
	}
	postParam.AuthorID = userID
	if postDto, err = p.postService.Create(&postParam); err != nil {
		ginx.ResponseError(c, e.CodeError)
		return
	}
	ginx.ResponseSuccess(c, postDto)
}
