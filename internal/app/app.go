package app

import (
	"fmt"
	"net/http"

	_ "xs.bbs/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"xs.bbs/internal/app/community"
	communityCtrl "xs.bbs/internal/app/community/controller"
	"xs.bbs/internal/app/user"
	userCtrl "xs.bbs/internal/app/user/controller"
	"xs.bbs/internal/pkg/middleware"
	"xs.bbs/pkg/conf"

	"github.com/gin-gonic/gin"
)

// Models gorm AutoMigrate 初始化
var Models = []interface{}{user.Model, community.Model}

// WebApp represent a web application
type WebApp struct {
	*gin.Engine
	Config        *conf.Config
	UserCtrl      *userCtrl.UserController
	CommunityCtrl *communityCtrl.CommunityController
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
	)
	r.GET("/ping", middleware.JWTAuth(), func(c *gin.Context) {
		c.String(http.StatusOK, "ping success")
	})
	return r, nil
}

// Start the web app
func (e *WebApp) Start() {
	e.Run(fmt.Sprintf(":%d", e.Config.System.Port))
}
