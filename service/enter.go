package service

import "myblog-server-go/service/system"

// @Title        enter.go
// @Description
// @Create       kurous 2025-08-21 13:44

var GroupApp = new(Group)

type Group struct {
	SystemServiceGroup system.ServiceGroup
}
