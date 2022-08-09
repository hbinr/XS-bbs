package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"

	"go.uber.org/zap"
	"xs.bbs/pkg/conf"
)

var rdb *redis.Client

// Init 初始化redis连接
func Init(cfg *conf.Config) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.RedisConfig.Host, cfg.RedisConfig.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default db
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleCons,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		zap.L().Error("redis ping failed", zap.Error(err))
		return nil, err
	}
	return rdb, nil
}

// Close 关闭redis client连接资源
func Close() {
	_ = rdb.Close()
}
