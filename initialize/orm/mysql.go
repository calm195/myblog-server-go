package orm

import (
	"myblog-server-go/config/orm"
	"myblog-server-go/global"
	"myblog-server-go/initialize/internal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @Title        mysql.go
// @Description
// @Create       kurous 2025-08-18 10:18

func GormMysql() *gorm.DB {
	mc := global.BlogConfig.Mysql
	return initGormMysql(mc)
}

func initGormMysql(mc orm.Mysql) *gorm.DB {
	if mc.Dbname == "" {
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       mc.Dsn(), // DSN data source name
		DefaultStringSize:         191,      // string 类型字段的默认长度
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}
	global.BlogLog.Infof("mysql config %v", mysqlConfig)

	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(mc.Prefix, mc.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+mc.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(mc.MaxIdleConnections)
		sqlDB.SetMaxOpenConns(mc.MaxOpenConnections)
		global.BlogLog.Infof("mysql connect success, database url: %s, dbname: %s", mc.Path, mc.Dbname)
		return db
	}
}
