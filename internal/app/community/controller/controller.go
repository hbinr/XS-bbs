package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"xs.bbs/internal/app/community/service"
)

var CommunityControllerSet = wire.NewSet(NewCommunityController)

type CommunityController struct {
	engine           *gin.Engine
	communityService service.ICommunityService
}

func NewCommunityController(e *gin.Engine, as service.ICommunityService) (*CommunityController, error) {
	community := &CommunityController{
		engine:           e,
		communityService: as,
	}
	g := e.Group("/api/community")
	{
		g.GET("/list", community.GetCommunityList)
	}
	return community, nil
}
