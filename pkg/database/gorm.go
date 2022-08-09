package database

import (
	"database/sql"

	"go.uber.org/zap"
	"xs.bbs/pkg/conf"

	"gorm.io/gorm/schema"

	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	sqlDB *sql.DB
)

// Init 初始化MySQL
func Init(cfg *conf.Config) (db *gorm.DB, err error) {
	mysqlConfig := mysql.Config{
		DSN:                       cfg.DSN, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	gormConfig := config(cfg.LogMode)
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		zap.L().Error("opens database failed", zap.Error(err))
		return
	}

	if sqlDB, err = db.DB(); err != nil {
		zap.L().Error("db.db() failed", zap.Error(err))
		return
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleCons)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenCons)
	return
}

// config 根据配置决定是否开启日志
func config(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	}
	return
}
