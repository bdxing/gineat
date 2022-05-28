package gineat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// Auto 添加自动注册路由
func (obj *RouterGroup) Auto(c interface{}) {
	autoHandle(c, func(relativePath string, hf func(*gin.Context)) {
		obj.Handle(http.MethodGet, relativePath, hf)
		obj.Handle(http.MethodPost, relativePath, hf)
	})
}
