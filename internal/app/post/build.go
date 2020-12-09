package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"xs.bbs/internal/app/post/controller"

	commuDao "xs.bbs/internal/app/community/dao"
	postDao "xs.bbs/internal/app/post/dao"
	"xs.bbs/internal/app/post/model"
	"xs.bbs/internal/app/post/service"
	userDao "xs.bbs/internal/app/user/dao"
)

var (
	Model = &model.Post{}
	//Set   = wire.NewSet(
	//	dao.PostDaoSet,
	//	service.PostServiceSet,
	//	controller.PostControllerSet,
	//)
)

func Init(engine *gin.Engine, db *gorm.DB, rdb *redis.Client) *controller.PostController {
	postDao := postDao.NewPostDao(db, rdb)
	userDao := userDao.NewUserDao(db)
	commuDao := commuDao.NewCommunityDao(db)
	postService := service.NewPostService(postDao, userDao, commuDao)
	return controller.NewPostController(engine, postService)
}
