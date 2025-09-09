package system

import "myblog-server-go/dao"

// @Title        enter.go
// @Description
// @Create       kurous 2025-08-21 13:45

type ServiceGroup struct {
	VisitService
}

var (
	visitDao = dao.GroupApp.SystemDaoGroup.VisitDao
)
