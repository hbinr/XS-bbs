package service

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	"github.com/gogf/gf/util/gconv"
	community "xs.bbs/internal/app/community/model"
	"xs.bbs/internal/app/post/model"
	post "xs.bbs/internal/app/post/model"
	user "xs.bbs/internal/app/user/model"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/util"
	"xs.bbs/pkg/utils/snowflake"
)

func (p *postService) Create(ctx context.Context, parm *post.PostParam) (err error) {
	var postModel post.Post
	if err = gconv.Struct(parm, &postModel); err != nil {
		err = errors.Wrap(e.ErrConvDataErr, err.Error())
		return
	}
	postModel.PostID = snowflake.GenID()
	return p.postRepo.Create(ctx, &postModel)
}

func (p *postService) GetPostByID(ctx context.Context, postID int64) (dto *post.PostDetailDto, err error) {
	var (
		postModel      *post.Post
		userModel      *user.User
		communityModel *community.Community
	)
	// 1.获取帖子
	if postModel, err = p.postRepo.GetPostByID(ctx, postID); err != nil {
		return
	}
	// 2.获取作者信息
	if userModel, err = p.userRepo.GetUserByID(ctx, postModel.AuthorID); err != nil {
		return
	}
	// 3.获取社区信息
	if communityModel, err = p.communityRepo.GetCommunityDetailByID(ctx, postModel.CommunityID); err != nil {
		return
	}
	if dto, err = ConvertToPostDetailDto(userModel, communityModel, postModel); err != nil {
		return
	}
	return
}

// GetPostListByIDs 根据post_id切片获取post列表，并按照给定的post_id顺序返回

func (p *postService) GetPostListByIDs(ctx context.Context, paging *util.PageInfo) (resList []*model.PostDetailDto, total int64, err error) {
	var (
		postListM  []*post.Post
		userM      *user.User
		communityM *community.Community
	)
	// 1.获取帖子列表
	if postListM, total, err = p.postRepo.GetPostList(ctx, paging.PageSize, paging.Offset()); err != nil {
		err = errors.Wrap(err, "service: GetPostList failed")
		return
	}
	resList = make([]*model.PostDetailDto, 0, len(postListM))
	for _, item := range postListM {
		// 2.获取作者信息
		if userM, err = p.userRepo.GetUserByID(ctx, item.AuthorID); err != nil {
			continue
		}
		// 3.获取社区信息
		if communityM, err = p.communityRepo.GetCommunityDetailByID(ctx, item.CommunityID); err != nil {
			continue
		}
		resDto := new(model.PostDetailDto)
		if resDto, err = ConvertToPostDetailDto(userM, communityM, item); err != nil {
			continue
		}
		resList = append(resList, resDto)
	}
	return
}

func ConvertToPostDetailDto(
	userM *user.User,
	communityM *community.Community,
	postM *post.Post) (dto *post.PostDetailDto, err error) {
	var (
		postDto      post.PostDto
		communityDto community.CommunityDto
	)
	if err = gconv.Struct(postM, &postDto); err != nil {
		err = errors.Wrap(e.ErrConvDataErr, err.Error())
		return
	}
	if err = gconv.Struct(communityM, &communityDto); err != nil {
		err = errors.Wrap(e.ErrConvDataErr, err.Error())
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

func (p *postService) Vote(ctx context.Context, userID int64, vote *model.PostVoteParam) (err error) {
	return p.postRepo.Vote(ctx, strconv.Itoa(int(userID)), strconv.Itoa(int(vote.PostID)), float64(vote.Direction))
}
