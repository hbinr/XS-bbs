package router

import (
	"XS-bbs/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	// 使用自定义封装的zap logger和recovery中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 测试
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test success...")
	})

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
