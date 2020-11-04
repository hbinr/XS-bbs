package dao

func (p *postDao) Create(post *PostModel) (err error) {
	return p.db.Create(post).Error
}
func (p *postDao) GetPostByID(pID int64) (post *PostModel, err error) {
	post = new(PostModel)
	err = p.db.Where("post_id", pID).First(&post).Error
	return
}

func (p *postDao) GetPostList(limit, offset int) (posts []*PostModel, total int64, err error) {
	posts = make([]*PostModel, 0, limit) // 默认取limit条
	db := p.db.Model(&PostModel{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&posts).Error
	return
}
