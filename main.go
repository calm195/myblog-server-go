package main

import (
	"myblog-server-go/core"
	"myblog-server-go/dao"
	"myblog-server-go/global"
	"myblog-server-go/initialize"
	"myblog-server-go/initialize/orm"
	"os"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	initializeSystem()
	core.RunServer()
}

func initializeSystem() {
	global.BlogVp = core.Viper() // 初始化Viper配置
	global.BlogLog = core.Zap()  // 初始化zap日志库
	global.BlogDb = orm.Gorm()   // 初始化Gorm数据库连接
	initialize.SetupHandlers()
	if global.BlogDb != nil {
		global.BlogLog.Info("registering db tables.")
		orm.RegisterTables() // 初始化数据库表
		dao.GroupApp.Init()
	} else {
		global.BlogLog.Errorf("registering db tables failed: %s", global.BlogConfig.System.DbType)
		os.Exit(0)
	}
}
