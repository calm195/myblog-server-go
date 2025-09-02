// Package global 公共包
package global

import (
	"gorm.io/gorm"
	"time"
)

// @Title        model.go
// @Description  公共模型，用于定义数据库表通用字段
// @Create       kurous 2025-08-17 14:16

// BlogModel
//
//	@Description: 定义数据库表通用字段
type BlogModel struct {
	ID        uint           `gorm:"primarykey" json:"-"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
