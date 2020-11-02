package dao

func (p *postDao) Create(post *PostModel) (err error) {
	post = new(PostModel)
	//if err = p.db.Create(&post).Error; err != nil {
	//	return err
	//}
	return p.db.Create(&post).Error
}
