package gineat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

type Controller struct {
	*Context
}

func (obj *Controller) Action() {
	obj.Ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func TestEat(t *testing.T) {
	r := Eat(gin.Default())

	gr := r.Group("home")

	gr.Auto(&Controller{})

	r.Run("0.0.0.0:8888")
}

func TestGin(t *testing.T) {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c.GetHeader("uid"))
		fmt.Println(c.GetHeader("rightLevel"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8888")
}
