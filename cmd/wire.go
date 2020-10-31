//+build wireinject

package main

import (
	"xs.bbs/internal/app"
	"xs.bbs/internal/app/community"
	"xs.bbs/internal/app/user"
	"xs.bbs/pkg/conf"
	"xs.bbs/pkg/database"

	"github.com/google/wire"
)

// initWebApp 注入函数，自定义的函数直接注入就行，不需要使用wire set
// 注意：log.Init()和snowflake.Init() 不需要提供provider，所以需要在生成wire_gen后手动加入它们的初始化
func initWebApp() (*app.WebApp, error) {
	// 逻辑顺序入参，未用到的依赖不需要注入
	wire.Build(
		conf.Init,                         // 初始化配置，自定义
		app.InitEngine,                    // 初始化web引擎，自定义
		database.Init,                     // 初始化mysql，自定义
		user.Set,                          // user业务provider，wire生成
		community.Set,                     // 获取业务controller
		wire.Struct(new(app.WebApp), "*"), // WebApp provider，wire生成
	)
	// 返回值不用管。直接返回nil就行
	return nil, nil
}
