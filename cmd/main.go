package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"

	"xs.bbs/internal/app"
	"xs.bbs/pkg/cache"
	"xs.bbs/pkg/conf"
	"xs.bbs/pkg/database"
	logger "xs.bbs/pkg/logger"
	"xs.bbs/pkg/utils/snowflake"
)

// @title 项目标题
// @version 0.0.1
// @description 项目描述
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8090
// @BasePath /api/
func main() {
	// 1. init config
	config, err := conf.Init()
	if err != nil {
		log.Fatalf("conf.Build failed, err: %+v", err)
	}

	// 2. init logger
	if err = logger.Init(config); err != nil {
		log.Fatalf("log.Build failed, err: %+v", err)
	}

	// 3. init snowflake
	if err = snowflake.Init(config); err != nil {
		zap.L().Error("snowflake.Build failed", zap.Error(err))
		return
	}

	// 4. init gorm client
	db, err := database.Init(config)
	if err != nil {
		zap.L().Error("database.Build failed", zap.Error(err))
		return
	}

	// 5. init gorm client
	rbd, err := cache.Init(config)
	if err != nil {
		zap.L().Error("cache.Build failed", zap.Error(err))
		return
	}

	if config.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	// 6. init app business
	router := app.Build(db, rbd)
	gracefulShutDown(router, config.System.Port)

}

func gracefulShutDown(router *gin.Engine, port int) {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// 7. start http server and gracefully shutdown
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
