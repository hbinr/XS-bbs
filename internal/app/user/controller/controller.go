package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"xs.bbs/internal/app/user/service"
)

var UserControllerSet = wire.NewSet(
	NewUserController)

type UserController struct {
	engine      *gin.Engine
	userService service.IUserService
}

func NewUserController(e *gin.Engine, us service.IUserService) (*UserController, error) {
	user := &UserController{
		engine:      e,
		userService: us,
	}
	g := e.Group("/api/user")
	{
		g.POST("/signup", user.SignUp)
		g.POST("/signin", user.SignIn)
		g.GET("/get", user.Get)
		g.GET("/delete", user.Delete)
	}
	return user, nil
}
