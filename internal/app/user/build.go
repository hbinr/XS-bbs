package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xs.bbs/internal/app/user/controller"
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/app/user/service"
)

var (
	Model = &model.User{}
	// 不使用wire依赖注入
	//Set   = wire.NewSet(
	//	dao.UserDaoSet,
	//	service.UserServiceSet,
	//	controller.UserControllerSet,
	//)
)

func Init(engine *gin.Engine, db *gorm.DB) *controller.UserController {
	userDao := dao.NewUserDao(db)
	userService := service.NewUserService(userDao)
	return controller.NewUserController(engine, userService)
}
