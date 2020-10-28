package cache

import (
	"fmt"

	"go.uber.org/zap"
	"xs.bbs/pkg/conf"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// Init 初始化redis连接
func Init(cfg *conf.RedisConfig) *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	if _, err := rdb.Ping().Result(); err != nil {
		zap.L().Error("redis ping failed", zap.Error(err))
	}
	return rdb
}

// RDB 获取redis客户端链接
func RDB() *redis.Client {
	return rdb
}

// Close 关闭redis clent连接资源
func Close() {
	_ = rdb.Close()
}
