package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xs.bbs/internal/app/user/controller"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/app/user/repository"
	"xs.bbs/internal/app/user/service"
)

var (
	Model = &model.User{}
	// 不使用wire依赖注入
	//Set   = wire.NewSet(
	//	repo.UserDaoSet,
	//	service.UserServiceSet,
	//	controller.UserControllerSet,
	//)
)

func Init(engine *gin.Engine, db *gorm.DB) *controller.UserController {
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	return controller.NewUserController(engine, userService)
}
