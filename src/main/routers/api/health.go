package api

import (
	"github.com/gin-gonic/gin"
	"star-im/src/main/common/app"
)

// Ping
// @Summary      健康检查
// @Description  接口连通性测试
// @Tags         测试
// @Success      200  {object}  app.Response
// @Router       /ping [get]
func Ping(c *gin.Context) {
	// 直接返回成功结果
	app.Success(c, nil)
}
