package model

// PostDto 帖子dto
type PostParam struct {
	AuthorID int64  `json:"authorID" form:"authorID"` // 作者的用户id
	Title    string `json:"title" form:"title"`       // 标题
	Content  string `json:"content" form:"content"`   // 内容
}
