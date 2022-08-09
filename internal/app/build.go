package app

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"xs.bbs/pkg/servers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	_ "xs.bbs/docs"
	"xs.bbs/internal/app/community"
	"xs.bbs/internal/app/post"
	"xs.bbs/internal/app/user"
)

var Entities = []interface{}{user.Entity, community.Entity, post.Entity}

func Build(db *gorm.DB, rdb *redis.Client) *gin.Engine {
	if err := db.AutoMigrate(Entities); err != nil {
		zap.L().Error("auto migrate  tables error", zap.Error(err))
	}

	return servers.NewHttpServer(
		user.Build(db),
		community.Build(db),
		post.Build(db, rdb),
	)
}
