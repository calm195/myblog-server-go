package post

import "myblog-server-go/global"

// @Title        post_category.go
// @Description
// @Create       kurous 2025-08-17 15:32

// PostCategory
//
//	@Description: 博客分类关联表
type PostCategory struct {
	global.BlogModel
	PostID     uint `json:"post_id" gorm:"not null;comment:文章ID"`     // 文章ID
	CategoryID uint `json:"category_id" gorm:"not null;comment:分类ID"` // 分类ID
}
