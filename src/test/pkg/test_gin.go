package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 测试 http 框架 https://github.com/gin-gonic/gin
func main() {
	// 默认返回一个已连接日志记录器和恢复中间件的引擎实例。
	r := gin.Default()
	// 绑定路由 /ping，访问后执行func的方法
	r.GET("/ping", func(c *gin.Context) {
		// 返回一个 json， 状态值为 200， H的内容为 map[string]
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 在0.0.0.0：8080上侦听和服务(对于Windows“为 localhost：8080”)
	err := r.Run()
	if err != nil {
		fmt.Println("启动服务异常：", err)
	}
}
