package controller

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/app/post/service"
)

//var PostControllerSet = wire.NewSet(
//	NewPostController)

type PostController struct {
	engine      *gin.Engine
	postService service.IPostService
}

func NewPostController(e *gin.Engine, us service.IPostService) *PostController {
	post := &PostController{
		engine:      e,
		postService: us,
	}

	g := e.Group("/api/post")

	{
		g.POST("/create", post.CreatePostHandle)
		g.GET("/post", post.GetPostDetailHandle)
		g.GET("/posts", post.GetPostListHandle)
		g.POST("/vote", post.VoteForPost)
	}

	return post
}
