package orm

import (
	"myblog-server-go/global"
	"myblog-server-go/model/post"
	"myblog-server-go/model/system"
	"os"

	"gorm.io/gorm"
)

// @Title        gorm.go
// @Description
// @Create       kurous 2025-08-18 10:31

func Gorm() *gorm.DB {
	currentDbType := global.BlogConfig.System.DbType
	if currentDbType == "" {
		global.BlogConfig.System.DbType = "mysql"
		currentDbType = "mysql"
	}
	global.BlogLog.Infof("current db type: %s", currentDbType)

	// 连接数据库
	// 目前支持 mysql 和 pgsql
	// 默认 mysql
	switch currentDbType {
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
		global.BlogLog.Errorf("register table failed: %s", err.Error())
		os.Exit(0)
	}

	global.BlogLog.Infof("register tables success")
}
