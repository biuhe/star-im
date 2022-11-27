package api

import (
	"github.com/gin-gonic/gin"
	"star-im/src/main/common/app"
)

// Ping 接口连通性测试
func Ping(c *gin.Context) {
	// 直接返回成功结果
	app.Success(c, nil)
}
