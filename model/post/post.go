package post

import "myblog-server-go/global"

// @Title        post.go
// @Description  技术博客
// @Create       kurous 2025-08-17 15:24

// Post
//
//	@Description: 博客表，包含标题、路径、内容字段
type Post struct {
	global.BlogModel
	Title   string `json:"title" gorm:"type:varchar(255);not null;comment:文章标题"`       // 文章标题
	Path    string `json:"path" gorm:"type:varchar(255);not null;unique;comment:文章路径"` // 文章路径
	Content string `json:"content" gorm:"type:text;not null;comment:文章内容"`             // 文章内容
}
