package system

import "myblog-server-go/service"

type ApiGroup struct {
	VisitApi
}

var (
	visitService = service.GroupApp.SystemServiceGroup.VisitService
)
