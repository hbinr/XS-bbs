package main

import (
	"XS-bbs/internal/app/router"
	"XS-bbs/pkg/db/mysql"
	"XS-bbs/pkg/db/redis"
	"XS-bbs/pkg/logger"
	"XS-bbs/pkg/setting"
	_ "XS-bbs/pkg/setting"
	"XS-bbs/pkg/util/snowflake"
	"fmt"
)

func main() {
	// 加载日志配置
	logger.Init(setting.Conf.LogConfig)
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 加载mysql配置
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接

	// 加载redis配置
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 初始化雪花算法配置
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
