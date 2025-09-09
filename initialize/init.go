package initialize

import (
	"myblog-server-go/global"
	"myblog-server-go/util"
)

// SetupHandlers
//
//	@Description: 注册系统重载处理函数
func SetupHandlers() {
	global.BlogLog.Info("setup system handlers...")

	// 注册系统重载处理函数
	global.BlogLog.Info("注册系统热重载函数")
	util.GlobalSystemEvents.RegisterReloadHandler(func() error {
		return util.Reload()
	})
}
