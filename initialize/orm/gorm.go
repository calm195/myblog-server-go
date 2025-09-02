package orm

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"myblog-server-go/global"
	"myblog-server-go/model/post"
	"myblog-server-go/model/system"
	"os"
)

// @Title        gorm.go
// @Description
// @Create       kurous 2025-08-18 10:31

func Gorm() *gorm.DB {
	switch global.BlogConfig.System.DbType {
	case "mysql":
		global.BlogActiveDbname = &global.BlogConfig.Mysql.Dbname
		return GormMysql()
	case "pgsql":
		global.BlogActiveDbname = &global.BlogConfig.Pgsql.Dbname
		return GormPgSql()
	default:
		global.BlogActiveDbname = &global.BlogConfig.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.BlogDb
	err := db.AutoMigrate(
		system.Visit{},
		post.Category{},
		post.Post{},
		post.PostCategory{},
	)
	if err != nil {
		global.BlogLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	global.BlogLog.Info("register table success")
}
