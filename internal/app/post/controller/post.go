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
		ginx.RespErrorWithMsg(c, e.CodeInvalidParams, errStr)
		return
	}

	if userID, err = ginx.GetCurrentUserID(c); err != nil {
		ginx.RespError(c, e.CodeNeedLogin)
		return
	}

	postParam.AuthorID = userID
	if err = p.postService.Create(&postParam); err != nil {
		ginx.RespError(c, e.CodeError)
		return
	}

	ginx.RespSuccess(c, nil)
}

func (p *PostController) GetPostDetailHandle(c *gin.Context) {
	var (
		pID int64
		err error
		dto *model.PostDetailDto
	)

	if pID, err = ginx.QueryInt("postID", c); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}

	if dto, err = p.postService.GetPostByID(pID); err != nil {
		ginx.RespErrorWithMsg(c, e.CodeError, err.Error())
		return
	}

	ginx.RespSuccess(c, dto)
}

func (p *PostController) GetPostListHandle(c *gin.Context) {
	var (
		err   error
		total int64
		posts []*model.PostDetailDto
	)

	pageInfo := common.PageInfo{
		Page:     1,
		PageSize: 5,
	}

	if errStr := ginx.BindAndValid(c, &pageInfo); errStr != "" {
		ginx.RespErrorWithMsg(c, e.CodeInvalidParams, errStr)
		return
	}

	if posts, total, err = p.postService.GetPostListByIDs(&pageInfo); err != nil {
		ginx.RespError(c, e.CodeError)
		return
	}

	pageRes := &common.PageResult{
		List:     posts,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}

	ginx.RespSuccess(c, pageRes)
}

func (p *PostController) VoteForPost(c *gin.Context) {
	var (
		err       error
		userID    int64
		voteParam model.PostVoteParam
	)

	if errStr := ginx.BindAndValid(c, &voteParam); errStr != "" {
		ginx.RespErrorWithMsg(c, e.CodeInvalidParams, errStr)
		return
	}

	if userID, err = ginx.GetCurrentUserID(c); err != nil {
		ginx.RespError(c, e.CodeError)
		return
	}

	if err = p.postService.Vote(userID, &voteParam); err != nil {
		ginx.RespError(c, e.CodeError)
		return
	}

	ginx.RespSuccess(c, nil)
}
