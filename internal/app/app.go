package app

import (
	"fmt"

	_ "xs.bbs/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"xs.bbs/internal/app/user/controller"
	"xs.bbs/internal/pkg/middleware"
	"xs.bbs/pkg/conf"

	"github.com/gin-gonic/gin"
)

// WebApp represent a web application
type WebApp struct {
	*gin.Engine
	Config   *conf.Config
	UserCtrl *controller.UserController
}

// InitEngine 初始化gin
func InitEngine(c *conf.Config) (*gin.Engine, error) {
	if c.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	// 设置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置公共中间件
	r.Use(
		middleware.GinLogger(),       // zap logger中间件
		middleware.GinRecovery(true), // zap recovery中间件
		middleware.Translation())     // 参数验证翻译中间件
	r.Group("/api")
	return r, nil
}

// Start the web app
func (e *WebApp) Start() {
	e.Run(fmt.Sprintf(":%d", e.Config.System.Port))
}
