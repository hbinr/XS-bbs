package main

import (
	"fmt"

	"go.uber.org/zap"

	"xs.bbs/internal/app"
	"xs.bbs/internal/app/community"
	"xs.bbs/internal/app/post"
	"xs.bbs/internal/app/user"

	"xs.bbs/pkg/cache"
	"xs.bbs/pkg/conf"
	"xs.bbs/pkg/database"
	"xs.bbs/pkg/log"
	"xs.bbs/pkg/tool/snowflake"
)

// @title 项目标题
// @version 0.0.1
// @description 这是一个gin web开发脚手架
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8090
// @BasePath /api/
func main() {
	webApp, err := initWebApp()
	if err != nil {
		panic(err)
	}
	webApp.Start()
}

func initWebApp() (webApp *app.WebApp, err error) {
	config, err := conf.Init()
	if err != nil {
		fmt.Println("conf.Init failed,err", err)
		return
	}
	if err = log.Init(config); err != nil {
		fmt.Println("log.Init failed,err", err)
		return
	}
	if err = snowflake.Init(config); err != nil {
		zap.L().Error("snowflake.Init failed", zap.Error(err))
		return
	}
	db, err := database.Init(config)
	if err != nil {
		zap.L().Error("database.Init failed", zap.Error(err))
		return
	}
	rbd, err := cache.Init(config)
	if err != nil {
		zap.L().Error("cache.Init failed", zap.Error(err))
		return
	}
	engine := app.InitEngine(config)
	webApp = &app.WebApp{
		Engine:        engine,
		Config:        config,
		UserCtrl:      user.Init(engine, db),
		CommunityCtrl: community.Init(engine, db),
		PostCtrl:      post.Init(engine, db, rbd),
	}
	return
}
