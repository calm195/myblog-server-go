package util

import (
	"sync"
)

// SystemEvents
//
//	@Description: 定义系统级事件处理
type SystemEvents struct {
	reloadHandlers []func() error
	mu             sync.RWMutex
}

// GlobalSystemEvents 全局事件管理器
var GlobalSystemEvents = &SystemEvents{}

// RegisterReloadHandler
//
//	@Description: 注册系统重载处理函数，将处理函数加到处理函数切片中，加锁添加
//	@receiver e
//	@param handler 处理函数
func (e *SystemEvents) RegisterReloadHandler(handler func() error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.reloadHandlers = append(e.reloadHandlers, handler)
}

// TriggerReload
//
//	@Description: 触发所有注册的重载处理函数 for循环执行切片中函数
//	@receiver e
//	@return error
func (e *SystemEvents) TriggerReload() error {
	e.mu.RLock()
	defer e.mu.RUnlock()

	for _, handler := range e.reloadHandlers {
		if err := handler(); err != nil {
			return err
		}
	}
	return nil
}
