package dao

import "myblog-server-go/dao/system"

// @Title        enter.go
// @Description
// @Create       kurous 2025-08-21 15:23

var GroupApp = new(Group)

type Group struct {
	SystemDaoGroup system.DaoGroup
}

func (g *Group) Init() {
	g.SystemDaoGroup.VisitDao.InitVisitDao()
}
