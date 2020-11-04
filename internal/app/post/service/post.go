package service

import (
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
	community "xs.bbs/internal/app/community/model"
	post "xs.bbs/internal/app/post/model"
	user "xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/util"
	"xs.bbs/pkg/tool/snowflake"
)

func (p *postService) Create(parm *post.PostParam) (err error) {
	var postModel post.Post
	if err = gconv.Struct(parm, &postModel); err != nil {
		zap.L().Error(e.CodeConvDataErr.Msg(), zap.Error(err))
		return
	}
	postModel.PostID = snowflake.GenID()
	return p.postDao.Create(&postModel)
}

func (p *postService) GetPostByID(pID int64) (dto *post.PostDetailDto, err error) {
	var (
		postModel      *post.Post
		userModel      *user.User
		communityModel *community.Community
	)
	// 1.获取帖子
	if postModel, err = p.postDao.GetPostByID(pID); err != nil {
		zap.L().Error("p.postDao.GetPostByID failed",
			zap.Int64("post_id", pID),
			zap.Error(err))
		return
	}
	// 2.获取作者信息
	if userModel, err = p.userDao.GetUserByID(postModel.AuthorID); err != nil {
		zap.L().Error("p.userDao.GetUserByID failed",
			zap.Int64("user_id", postModel.AuthorID),
			zap.Error(err))
		return
	}
	// 3.获取社区信息
	if communityModel, err = p.communityDao.GetCommunityDetailByID(postModel.CommunityID); err != nil {
		zap.L().Error("p.communityDao.GetCommunityDetailByID failed",
			zap.Int64("community_id", communityModel.CommunityID),
			zap.Error(err))
		return
	}
	if err = gconv.Struct(postModel, &dto); err != nil {
		zap.L().Error(e.CodeConvDataErr.Msg(), zap.Error(err))
		return
	}
	if dto, err = Convert2PostDetailDto(userModel, communityModel, postModel); err != nil {
		return
	}
	return
}

func Convert2PostDetailDto(userM *user.User, communityM *community.Community,
	postM *post.Post) (dto *post.PostDetailDto, err error) {
	var (
		postDto      post.PostDto
		communityDto community.CommunityDto
	)
	if err = gconv.Struct(postM, &postDto); err != nil {
		return nil, err
	}
	if err = gconv.Struct(communityM, &communityDto); err != nil {
		return nil, err
	}
	dto = &post.PostDetailDto{
		UserName:     userM.Username,
		PostDto:      &postDto,
		CommunityDto: &communityDto,
	}
	dto.CreatedAt = util.TimeFormat(postM.CreatedAt, util.FMT_DATE_TIME)
	return
}
