package post

import "myblog-server-go/global"

// @Title        category.go
// @Description
// @Create       kurous 2025-08-17 15:31

// Category
//
//	@Description: 分类表，包含分类名称字段
type Category struct {
	global.BlogModel
	Name string `json:"name" gorm:"type:varchar(100);not null;unique;comment:分类名称"` // 分类名称
}
