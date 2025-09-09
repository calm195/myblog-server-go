package system

import "myblog-server-go/api"

// @Title        enter.go
// @Description
// @Create       kurous 2025-08-18 17:38

type RouterGroup struct {
	VisitRouter
}

var (
	visitApi = api.GroupApp.SystemApiGroup.VisitApi
)
