package gineat

import (
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
}

// Auto 添加自动注册路由
func (obj *Engine) Auto(httpMethod string, c interface{}) {
	autoHandle(c, func(relativePath string, hf func(*gin.Context)) {
		obj.Handle(httpMethod, relativePath, hf)
	})
}

// Group 包装分组
func (obj *Engine) Group(relativePath string) *RouterGroup {
	return &RouterGroup{obj.RouterGroup.Group(relativePath)}
}
