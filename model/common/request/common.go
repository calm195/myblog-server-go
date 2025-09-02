package request

import (
	"gorm.io/gorm"
)

// PageInfo
//
// @Description: Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   // 关键字
}

// Paginate Gorm 分页函数
//
//	@Description: 默认每页8条数据，最大32条数据，起始页码为1。by db.Offset(offset).Limit(pageSize)
//	@receiver r 分页参数
//	@return func(db *orm.DB) *orm.DB
func (r *PageInfo) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page <= 0 {
			r.Page = 1
		}
		switch {
		case r.PageSize > 32:
			r.PageSize = 32
		case r.PageSize <= 0:
			r.PageSize = 8
		}
		offset := (r.Page - 1) * r.PageSize
		return db.Offset(offset).Limit(r.PageSize)
	}
}

// Empty 用于表示空请求体
type Empty struct{}
