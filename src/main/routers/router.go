package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
	_ "star-im/src/main/docs"
	"star-im/src/main/handler"
	"star-im/src/main/routers/api"
)

// Setup 初始化路由
func Setup() *gin.Engine {
	r := gin.Default()
	// 记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 使用中间件
	// 统一异常处理
	r.Use(handler.Recover)
	// 统一日志
	r.Use(gin.Logger())

	// 不需要鉴权
	r.GET("/ping", api.Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
