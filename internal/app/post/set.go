package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"xs.bbs/internal/app/post/controller"
	"xs.bbs/internal/app/post/model"
	"xs.bbs/internal/app/post/service"
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
	cs := service.NewPostService(db, rdb)
	return controller.NewPostController(engine, cs)
}
