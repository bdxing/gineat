package gineat

import "github.com/gin-gonic/gin"

func Eat(e *gin.Engine) *Engine {
	return &Engine{e}
}
