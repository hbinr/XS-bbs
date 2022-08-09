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

func NewPostController(us service.PostService) *PostController {
	return &PostController{
		postService: us,
	}
}

func (p *PostController) RegisterHTTPRouter(r *gin.Engine) {
	g := r.Group("/api/post")

	{
		g.POST("/", p.CreatePostHandle)
		g.GET("/info", p.GetPostDetailHandle)
		g.GET("/list", p.GetPostListHandle)
		g.POST("/vote", p.VoteForPost)
	}

}
