package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/post/model"
	"xs.bbs/internal/pkg/common"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/ginx"
)

func (p *PostController) CreatePostHandle(c *gin.Context) {
	var (
		err       error
		userID    int64
		postParam model.PostParam
	)
	if errStr := ginx.BindAndValid(c, &postParam); errStr != "" {
		ginx.ResponseErrorWithMsg(c, e.CodeInvalidParams, errStr)
		return
	}
	if userID, err = ginx.GetCurrentUserID(c); err != nil {
		ginx.ResponseError(c, e.CodeNeedLogin)
		return
	}
	postParam.AuthorID = userID
	if err = p.postService.Create(&postParam); err != nil {
		ginx.ResponseError(c, e.CodeError)
		return
	}
	ginx.ResponseSuccess(c, nil)
}

func (p *PostController) GetPostDetailHandle(c *gin.Context) {
	var (
		pID int64
		err error
		dto *model.PostDetailDto
	)

	if pID, err = ginx.QueryInt("postID", c); err != nil {
		ginx.ResponseError(c, e.CodeInvalidParams)
		return
	}
	if dto, err = p.postService.GetPostByID(pID); err != nil {
		ginx.ResponseError(c, e.CodeError)
		return
	}
	ginx.ResponseSuccess(c, dto)
}

func (p *PostController) GetPostListHandle(c *gin.Context) {
	var (
		err      error
		total    int64
		pageInfo common.PageInfo
		posts    []*model.PostDetailDto
	)
	if errStr := ginx.BindAndValid(c, &pageInfo); errStr != "" {
		ginx.ResponseErrorWithMsg(c, e.CodeInvalidParams, errStr)
		return
	}
	if posts, total, err = p.postService.GetPostList(&pageInfo); err != nil {
		ginx.ResponseError(c, e.CodeError)
		return
	}
	pageRes := &common.PageResult{
		List:     posts,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}
	ginx.ResponseSuccess(c, pageRes)
}
