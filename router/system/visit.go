package system

import "github.com/gin-gonic/gin"

// @Title        visit.go
// @Description
// @Create       kurous 2025-08-21 10:58

type VisitRouter struct{}

func (v *VisitRouter) InitVisitRouter(Router *gin.RouterGroup) {
	visitRouter := Router.Group("visit")

	visitRouter.GET("getVisitCount", visitApi.GetVisitCount) // 获取访问量
}
