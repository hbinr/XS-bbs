package post

import (
	"github.com/gin-gonic/gin"
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

func Init(engine *gin.Engine, db *gorm.DB) *controller.PostController {
	cs := service.NewPostService(db)
	return controller.NewPostController(engine, cs)
}
