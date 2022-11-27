package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"star-im/src/main/common/app"
)

// Recover 注意 Recover 要尽量放在router.User的第一个被加载
// 如不是的话，在recover前的中间件或路由，将不能被拦截到
// 程序的原理是：
// 1.请求进来，执行recover
// 2.程序异常，抛出panic
// 3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
func Recover(c *gin.Context) {
	defer func() {
		r := recover()
		if r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			c.JSON(http.StatusOK, app.Response{
				Code: app.ERROR,
				Msg:  ErrorToString(r),
				Data: nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()

	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// ErrorToString recover错误，转string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
