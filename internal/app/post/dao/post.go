package dao

func (p *postDao) Create(post *PostModel) (err error) {
	return p.db.Create(post).Error
}
func (p *postDao) GetPostByID(pID int64) (post *PostModel, err error) {
	post = new(PostModel)
	err = p.db.Where("post_id", pID).First(&post).Error
	return
}
