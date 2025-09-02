package system

import (
	"myblog-server-go/global"
	"myblog-server-go/model/system"
)

// @Title        visit.go
// @Description
// @Create       kurous 2025-08-21 15:23

type VisitDao struct{}

var db = global.BlogDb.Model(&system.Visit{})

func (v *VisitDao) GetVisitCount() (count int64) {
	db.Count(&count)
	return
}
