package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/community/service"
	"xs.bbs/internal/pkg/middleware"
)

//var CommunityControllerSet = wire.NewSet(NewCommunityController)

type CommunityController struct {
	engine           *gin.Engine
	communityService service.ICommunityService
}

func NewCommunityController(e *gin.Engine, as service.ICommunityService) *CommunityController {
	community := &CommunityController{
		engine:           e,
		communityService: as,
	}

	e.Use(middleware.JWTAuth())
	g := e.Group("/api/community")

	{
		g.GET("/communities", community.GetCommunityList)
		g.GET("/detail", community.GetCommunityDetail)
	}

	return community
}
