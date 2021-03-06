package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/user/service"
)

//var UserControllerSet = wire.NewSet(
//	NewUserController)

type UserController struct {
	engine      *gin.Engine
	userService service.IUserService
}

func NewUserController(e *gin.Engine, us service.IUserService) *UserController {
	user := &UserController{
		engine:      e,
		userService: us,
	}

	g := e.Group("/api/user")

	{
		g.POST("/signup", user.SignUp)
		g.POST("/signin", user.SignIn)
		g.GET("/:userID", user.Get)
		g.DELETE("/:userID", user.Delete)
	}

	return user
}
