package util

import (
	"go.uber.org/zap"
	"myblog-server-go/global"
	"myblog-server-go/initialize/orm"
)

// Reload 优雅地重新加载系统配置
func Reload() error {
	global.BlogLog.Info("正在重新加载系统配置...")

	// 重新加载配置文件
	if err := global.BlogVp.ReadInConfig(); err != nil {
		global.BlogLog.Error("重新读取配置文件失败!", zap.Error(err))
		return err
	}

	// 重新初始化数据库连接
	if global.BlogDb != nil {
		db, _ := global.BlogDb.DB()
		err := db.Close()
		if err != nil {
			global.BlogLog.Error("关闭原数据库连接失败!", zap.Error(err))
			return err
		}
	}

	// 重新建立数据库连接
	global.BlogDb = orm.Gorm()

	if global.BlogDb != nil {
		// 确保数据库表结构是最新的
		orm.RegisterTables()
	}

	global.BlogLog.Info("系统配置重新加载完成")
	return nil
}
