// Package global 公共包
package global

import (
	"time"

	"gorm.io/gorm"
)

// @Title        model.go
// @Description  公共模型，用于定义数据库表通用字段
// @Create       kurous 2025-08-17 14:16

// BlogModel
//
//	@Description: 定义数据库表通用字段
type BlogModel struct {
	ID        uint           `gorm:"primarykey" json:"-"` // 主键ID
	CreatedAt time.Time      `json:"created_at"`          // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`          // 更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"`          // 删除时间
}
