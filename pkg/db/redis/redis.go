package redis

import (
	"XS-bbs/pkg/setting"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
)

// Init 初始化redis连接
func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	_, err = rdb.Ping().Result()
	return
}

// Close 关闭redis clent连接资源
func Close() {
	_ = rdb.Close()
}
