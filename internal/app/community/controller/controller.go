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

func NewCommunityController(e *gin.Engine, as service.CommunityService) *CommunityController {
	community := &CommunityController{
		engine:           e,
		communityService: as,
	}

	e.Use(middleware.JWTAuth())
	g := e.Group("/api/community")

	{
		g.GET("/list", community.GetCommunityList)
		g.GET("/info", community.GetCommunityDetail)
	}

	return community
}
