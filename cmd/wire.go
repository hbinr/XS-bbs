//+build wireinject

package main

import (
	"xs.bbs/internal/app"
	"xs.bbs/internal/app/user"
	"xs.bbs/internal/app/user/controller"
	"xs.bbs/pkg/conf"
	"xs.bbs/pkg/database"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewUseController)

// initWebApp 注入函数，自定义的函数直接注入就行，不需要使用wire set
func initWebApp() (*app.WebApp, error) {
	// 逻辑顺序入参，未用到的依赖不需要注入
	wire.Build(
		conf.Init,                         // 初始化配置，自定义
		app.InitEngine,                    // 初始化web引擎，自定义
		database.Init,                     // 初始化mysql，自定义
		user.Set,                          // user业务provider，wire生成
		controllerSet,                     // 获取业务controller
		wire.Struct(new(app.WebApp), "*"), // WebApp provider，wire生成
	)
	// 返回值不用管。直接返回nil就行
	return nil, nil
}
