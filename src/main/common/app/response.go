package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应对象
type Response struct {
	// 响应编码
	Code int `json:"code"`
	// 返回消息
	Msg string `json:"msg"`
	// 返回数据
	Data interface{} `json:"data"`
}

// Res 设置 gin.JSON 的内容
func Res(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}

// Success 返回成功结果
func Success(c *gin.Context, data interface{}) {
	Res(c, http.StatusOK, SUCCESS, data)
}

// Error 返回错误结果，异常结果放在统一异常处理 handler中
func Error(c *gin.Context, data interface{}) {
	Res(c, http.StatusOK, ERROR, data)
}
