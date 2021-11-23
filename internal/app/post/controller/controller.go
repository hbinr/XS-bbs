package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/post/service"
)

//var PostControllerSet = wire.NewSet(
//	NewPostController)

type PostController struct {
	engine      *gin.Engine
	postService service.PostService
}

func NewPostController(e *gin.Engine, us service.PostService) *PostController {
	post := &PostController{
		engine:      e,
		postService: us,
	}

	g := e.Group("/api/post")

	{
		g.POST("/", post.CreatePostHandle)
		g.GET("/info", post.GetPostDetailHandle)
		g.GET("/list", post.GetPostListHandle)
		g.POST("/vote", post.VoteForPost)
	}

	return post
}
