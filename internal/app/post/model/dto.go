package model

import "xs.bbs/internal/app/community/model"

// PostDto 帖子dto
type PostDto struct {
	PostID      int64  `json:"postID" form:"postID"`           // 帖子id
	AuthorID    int64  `json:"authorID" form:"authorID"`       // 作者的用户id
	CommunityID int64  `json:"communityID" form:"communityID"` // 所属社区
	Status      int8   `json:"status" form:"status"`           // 帖子状态 1:有效,0:无效
	Title       string `json:"title" form:"title"`             // 标题
	Content     string `json:"content" form:"content"`         // 内容
	CreatedAt   string `json:"createdAt" form:"createdAt"`     // 创建时间
}

// PostDetailDto 帖子详情dto
type PostDetailDto struct {
	UserName string `json:"userName"`
	*PostDto
	*model.CommunityDto `json:"communityDto"`
}
