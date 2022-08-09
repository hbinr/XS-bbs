package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/community/service"
	"xs.bbs/internal/pkg/middleware"
)

//var CommunityControllerSet = wire.NewSet(NewCommunityController)

type CommunityController struct {
	engine           *gin.Engine
	communityService service.CommunityService
}

func NewCommunityController(service service.CommunityService) *CommunityController {
	return &CommunityController{
		communityService: service,
	}
}

func (cc *CommunityController) RegisterHTTPRouter(r *gin.Engine) {
	r.Use(middleware.JWTAuth())
	g := r.Group("/api/community")

	{
		g.GET("/list", cc.GetCommunityList)
		g.GET("/info", cc.GetCommunityDetail)
	}

	return
}
