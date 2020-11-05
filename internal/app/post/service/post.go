package service

import (
	"strconv"

	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
	community "xs.bbs/internal/app/community/model"
	"xs.bbs/internal/app/post/model"
	post "xs.bbs/internal/app/post/model"
	user "xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/common"
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
	if dto, err = Convert2PostDetailDto(userModel, communityModel, postModel); err != nil {
		return
	}
	return
}

func (p *postService) GetPostList(paging *common.PageInfo) (resList []*model.PostDetailDto, total int64, err error) {
	var (
		postListM  []*post.Post
		userM      *user.User
		communityM *community.Community
	)
	// 1.获取帖子列表
	if postListM, total, err = p.postDao.GetPostList(paging.PageSize, paging.Offset()); err != nil {
		zap.L().Error("p.userDao.GetPostList failed", zap.Error(err))
		return
	}
	resList = make([]*model.PostDetailDto, 0, len(postListM))
	for _, post := range postListM {
		// 2.获取作者信息
		if userM, err = p.userDao.GetUserByID(post.AuthorID); err != nil {
			zap.L().Error("p.userDao.GetUserByID failed",
				zap.Int64("user_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 3.获取社区信息
		if communityM, err = p.communityDao.GetCommunityDetailByID(post.CommunityID); err != nil {
			zap.L().Error("p.communityDao.GetCommunityDetailByID failed",
				zap.Int64("community_id", communityM.CommunityID),
				zap.Error(err))
			continue
		}
		resDto := new(model.PostDetailDto)
		if resDto, err = Convert2PostDetailDto(userM, communityM, post); err != nil {
			zap.L().Error("post.service.Convert2PostDetailDto failed", zap.Error(err))
			continue
		}
		resList = append(resList, resDto)
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
		return
	}
	if err = gconv.Struct(communityM, &communityDto); err != nil {
		return
	}
	dto = &post.PostDetailDto{
		UserName:     userM.Username,
		PostDto:      &postDto,
		CommunityDto: &communityDto,
	}
	dto.CreatedAt = util.TimeFormat(postM.CreatedAt, util.FMT_DATE_TIME)
	return
}

func (p *postService) Vote(userID int64, vote *model.PostVoteParam) (err error) {
	zap.L().Debug("Vote",
		zap.Int64("userID", userID),
		zap.Int64("postID", vote.PostID),
		zap.Int8("direction", vote.Direction))
	return p.postDao.Vote(strconv.Itoa(int(userID)), strconv.Itoa(int(vote.PostID)), float64(vote.Direction))
}
