package gineat

import (
	"fmt"
	"github.com/bdxing/goutils"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

// exceptMethod 过滤控制器不反射的方法
var exceptMethod = []string{"ResponseCtx", "ResponsePages"}

// autoHandle 自动控制器反射路由
func autoHandle(c interface{}, handlerFuncCall func(string, func(*gin.Context))) {
	reflectVal := reflect.ValueOf(c)

	rt := reflectVal.Type()
	ct := reflect.Indirect(reflectVal).Type()

	for i := 0; i < rt.NumMethod(); i++ {
		if goutils.SliceInStr(rt.Method(i).Name, exceptMethod) {
			continue
		}

		handlerFunc := func(ctx *gin.Context) {

			paths := strings.Split(ctx.FullPath(), "/")

			methodName := goutils.StrUnderToHump(paths[len(paths)-1])

			nrt := reflect.New(ct)
			nct := reflect.Indirect(nrt).Type()

			if _, b := nct.FieldByName("Context"); b {
				nrt.Elem().FieldByName("Context").Set(reflect.ValueOf(&Context{ctx}))
			}

			nrt.MethodByName(methodName).Call([]reflect.Value{})
		}

		relativePath := fmt.Sprintf("/%s/%s", goutils.StrHumpToUnder(ct.Name()), goutils.StrHumpToUnder(rt.Method(i).Name))

		handlerFuncCall(relativePath, handlerFunc)
	}
}
