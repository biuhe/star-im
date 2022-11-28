package main

import (
	"fmt"
	"log"
	"net/http"
	"star-im/src/main/config"
	"star-im/src/main/config/settings"
	"star-im/src/main/routers"
)

// init 初始化
func init() {
	// 初始化配置项
	config.Initial()
}

// @title           Star-IM
// @version         1.0
// @description     即时通讯接口文档
// #termsOfService  http://swagger.io/terms/

// @contact.name   Biu He
// @contact.url    http://www.swagger.io/support
// @contact.email  wsxc_0617@sina.cn

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// BasePath  /
// @schemes http https
// #securityDefinitions.basic  BasicAuth
func main() {
	// 路由
	routersInit := routers.Setup()
	// 读取超时
	readTimeout := settings.ServerSetting.ReadTimeout
	// 写入超时
	writeTimeout := settings.ServerSetting.WriteTimeout
	// 端口
	endPoint := fmt.Sprintf(":%d", settings.ServerSetting.Port)
	// 最大 header 数
	maxHeaderBytes := 1 << 20

	// 配置 http Server
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] 启动http服务器侦听 %s", endPoint)

	// 启动服务
	err := server.ListenAndServe()
	if err != nil {
		// 启动异常
		panic(fmt.Errorf("启动服务异常：%w", err))
	}
}
