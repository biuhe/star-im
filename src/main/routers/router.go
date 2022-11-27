package routers

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
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

	// swagger info
	//docs.SwaggerInfo.Title = "Star-Im"
	//docs.SwaggerInfo.Description = "即时通讯接口文档"
	//docs.SwaggerInfo.Version = "1.0"
	////docs.SwaggerInfo.Host = "petstore.swagger.io"
	////docs.SwaggerInfo.BasePath = "/v2"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
