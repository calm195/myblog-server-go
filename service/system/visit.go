package system

import "myblog-server-go/model/system"

// @Title        enter.go
// @Description
// @Create       kurous 2025-08-21 13:45

type VisitService struct{}

func (v *VisitService) AddVisit(visit *system.Visit) (err error) {
	err = visitDao.Save(*visit)
	return
}

func (v *VisitService) GetVisitList() (list []system.Visit) {
	list = visitDao.GetVisitList()
	return
}

func (v *VisitService) GetVisitCount() (count int64) {
	count = visitDao.GetVisitCount()
	return
}
