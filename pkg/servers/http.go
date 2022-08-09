package servers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"xs.bbs/internal/pkg/middleware"
)

// Server http  server
type Server interface {
	// RegisterHTTPRouter register http router
	RegisterHTTPRouter(r *gin.Engine)
}

func NewHttpServer(servers ...Server) *gin.Engine {

	r := gin.New()
	// 设置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置公共中间件
	r.Use(
		middleware.GinLogger(),       // zap logger中间件
		middleware.GinRecovery(true), // zap recovery中间件
	)

	for _, s := range servers {
		s.RegisterHTTPRouter(r)
	}

	return r

}
