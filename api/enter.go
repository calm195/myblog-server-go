package api

import "myblog-server-go/api/system"

// @Title        enter.go
// @Description
// @Create       kurous 2025-08-18 17:40

var GroupApp = new(Group)

type Group struct {
	SystemApiGroup system.ApiGroup
}
