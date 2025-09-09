package router

import "myblog-server-go/router/system"

// @Title        enter.go
// @Description  路由入口
// @Create       kurous 2025-08-17 14:54

var GroupApp = new(Group)

type Group struct {
	// InitRouter 初始化路由
	SystemRouterGroup system.RouterGroup
}
