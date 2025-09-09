package system

import (
	"myblog-server-go/model/common/response"

	"github.com/gin-gonic/gin"
)

type VisitApi struct{}

func (v *VisitApi) GetVisitCount(c *gin.Context) {
	count := visitService.GetVisitCount()
	response.OkWithData(count, c)
}
