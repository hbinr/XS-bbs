package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/user/service"
	"xs.bbs/internal/pkg/middleware"
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
	g.Use(middleware.JWT) // 设置user私有中间件
	{
		g.POST("/signup", user.SignUp)
		g.POST("/signin", user.SignIn)
		g.GET("/get", user.Get)
		g.GET("/delete", user.Delete)
	}
	return user, nil
}
