package community

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xs.bbs/internal/app/community/controller"
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/app/community/service"
)

var (
	Model = &model.Community{}
	//Set   = wire.NewSet(
	//	dao.CommunityDaoSet,
	//	service.CommunityServiceSet,
	//	controller.CommunityControllerSet,
	//)
)

func Init(engine *gin.Engine, db *gorm.DB) *controller.CommunityController {
	cs := service.NewCommunityService(db)
	return controller.NewCommunityController(engine, cs)
}
