package model

// PostParam 帖子dto
type PostParam struct {
	AuthorID int64  `json:"authorID" form:"authorID"` // 作者的用户id
	Title    string `json:"title" form:"title"`       // 标题
	Content  string `json:"content" form:"content"`   // 内容
}

// PostVoteParam 帖子投票dto
type PostVoteParam struct {
	PostID    int64 `json:"postID" form:"postID" v:"postID@required|min:1 #请输入帖子ID|请输入帖子ID"`             //  帖子ID;min:1用来保证postID必须传入                  //
	Direction int8  `json:"direction" form:"direction" v:"direction@in:1,-1,0#direction参数值应该在(1,-1,0)中"` // // 赞成票:1;反对票:-1;取消投票:0
}
