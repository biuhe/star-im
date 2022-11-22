// main 方法必须使用 main 包
package main

// 引入依赖
import (
	"fmt"
	"github.com/spf13/viper"
)

// 主要执行的方法
// 测试 viper 配置管理库 https://github.com/spf13/viper
func main() {
	// 配置文件名(不带扩展名，即 app.yml 只需要app这部分)
	viper.SetConfigName("app")
	// 如果配置文件名称中没有扩展名，则为必填项
	viper.SetConfigType("yaml")
	// 在其中查找配置文件的路径
	viper.AddConfigPath("src/resource/")
	// 查找并读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		// 处理读取配置文件时出现的错误
		panic(fmt.Errorf("读取配置异常，原因为: %w", err))
	}
	// 打印内容到控制台
	fmt.Println("初始化 app 配置成功")
	// 获取配置文件中的参数
	url := viper.GetString("settings.server.url")
	port := viper.GetString("settings.server.port")
	// 打印参数
	fmt.Printf("配置中的服务器地址及端口号为：%s:%s", url, port)
}
