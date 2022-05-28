## <p align="center">gineat</p>
This is a toolkit that enables the GIN framework to support automatic routing.


## Installation

To install this package, you need to setup your Go workspace.  The simplest way to install the library is to run:

```
go get github.com/bdxing/gineat
```


## Example 

```go
package main

import (
	"github.com/bdxing/gineat"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gineat.Context
}

func (obj *Controller) Action() {
	obj.Ctx.JSON(200, gin.H{
		"message": "hello word",
	})
}

func main() {
	r := gineat.Eat(gin.Default())

	r.Auto(&Controller{})

	r.Run("0.0.0.0:8888")

	// request url: http://localhost:8888/controller/action
}

```