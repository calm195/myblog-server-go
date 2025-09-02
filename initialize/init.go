package initialize

import (
	"myblog-server-go/util"
)

// 初始化全局函数
func SetupHandlers() {
	// 注册系统重载处理函数
	util.GlobalSystemEvents.RegisterReloadHandler(func() error {
		return util.Reload()
	})
}
