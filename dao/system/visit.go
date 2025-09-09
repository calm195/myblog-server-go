package system

import (
	"math"
	"myblog-server-go/global"
	"myblog-server-go/model/system"

	"gorm.io/gorm"
)

// @Title        enter.go
// @Description  system.Visit Dao层
// @Create       kurous 2025-08-21 15:23

type IVisitDao interface {
	Save(visit system.Visit) error
	GetVisitCount() (count int64)
	GetTotalPageCount() (count int)
	GetVisitList() (lists []system.Visit)
	GetVisitListByPage(page int, pageSize int) (lists []system.Visit)
}

type VisitDao struct {
	db *gorm.DB
}

func (v *VisitDao) InitVisitDao() {
	v.db = global.BlogDb.Model(&system.Visit{})
}

// Save
//
//	@Description: 保存访问记录
//	@receiver v
//	@param visit
//	@return error
func (v *VisitDao) Save(visit system.Visit) error {
	return v.db.Create(&visit).Error
}

// GetVisitCount
//
//	@Description: 获取访问记录总数
//	@receiver v
//	@return count
func (v *VisitDao) GetVisitCount() (count int64) {
	global.BlogDb.Model(&system.Visit{}).Count(&count)
	return
}

// GetTotalPageCount
//
//	@Description: 获取总页数
//	@receiver v
//	@return count
func (v *VisitDao) GetTotalPageCount() (count int) {
	totalCount := v.GetVisitCount()
	count = int(math.Ceil(float64(totalCount) / float64(100)))
	return
}

// GetVisitList
//
//	@Description: 获取所有访问记录
//	@receiver v
//	@return lists
func (v *VisitDao) GetVisitList() (lists []system.Visit) {
	global.BlogDb.Model(&system.Visit{}).Select(system.Visit{}).Find(&lists)
	return
}

// GetVisitListByPage
//
//	@Description: 分页查询所有访问记录，如果传入参数不合法，那么采用默认分页
//	@receiver v
//	@param page 第几页，默认为1
//	@param pageSize 每页大小，默认为10
//	@return lists
func (v *VisitDao) GetVisitListByPage(page int, pageSize int) (lists []system.Visit) {
	if pageSize == 0 || pageSize > 100 || page < 1 {
		page = 1
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	global.BlogDb.Model(&system.Visit{}).Limit(pageSize).Offset(offset).Find(&lists)
	return
}
