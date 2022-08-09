package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/user/service"
)

//var UserControllerSet = wire.NewSet(
//	NewUserController)

type UserController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (u *UserController) RegisterHTTPRouter(r *gin.Engine) {
	g := r.Group("/api/user")

	{
		g.POST("/signup", u.Register)
		g.POST("/signin", u.Login)
		g.GET("/:userID", u.Get)
		g.DELETE("/:userID", u.Delete)
	}
}
