package system

import "myblog-server-go/global"

// @Title        enter.go
// @Description
// @Create       kurous 2025-08-17 15:36

type Visit struct {
	global.BlogModel
	IP   string `json:"ip" gorm:"type:varchar(45);not null;comment:访问IP"`    // 访问IP
	Path string `json:"path" gorm:"type:varchar(255);not null;comment:访问路径"` // 访问路径
}
