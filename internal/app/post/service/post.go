package service

import (
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
	"xs.bbs/internal/app/post/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/pkg/tool/snowflake"
)

func (p *postService) Create(parm *model.PostParam) (dto *PostDto, err error) {
	var postModel model.Post
	if err = gconv.Struct(parm, &postModel); err != nil {
		zap.L().Error(e.CodeConvDataErr.Msg(), zap.Error(err))
		return
	}
	postModel.PostID = snowflake.GenID()
	if err = p.dao.Create(&postModel); err != nil {
		zap.L().Error("postService.Create failed ", zap.Error(err))
		return
	}
	if err = gconv.Struct(postModel, &dto); err != nil {
		zap.L().Error(e.CodeConvDataErr.Msg(), zap.Error(err))
		return
	}
	return
}
