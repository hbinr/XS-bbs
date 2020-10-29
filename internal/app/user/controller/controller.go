package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/user/service"
)

type UserController struct {
	engine      *gin.Engine
	userService service.IUserService
}

func NewUseController(e *gin.Engine, us service.IUserService) (*UserController, error) {
	user := &UserController{
		engine:      e,
		userService: us,
	}
	g := e.Group("/user")
	{
		g.POST("/signup", user.SignUp)
		g.POST("/signin", user.SignIn)
		g.GET("/get", user.Get)
		g.GET("/delete", user.Delete)
	}
	return user, nil
}
