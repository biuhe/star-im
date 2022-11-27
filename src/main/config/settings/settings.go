// Package settings
/**
此文件用于读取配置文件 app.yml，并设置到对应实体，以提供给其他文件使用。
因此该文件需要优先进行初始化操作
*/
package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// 定义实体装载配置文件内容

// Settings 设置
type Settings struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"db"`
}

var Setting = &Settings{}

// Server 服务
type Server struct {
	// Url 地址
	Url string `mapstructure:"url"`
	// Port 端口
	Port int `mapstructure:"port"`
	// ReadTimeout 读取超时
	ReadTimeout time.Duration `mapstructure:"readTimeout"`
	// WriteTimeout 写入超时
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}

var ServerSetting = &Server{}

// Database 数据库
type Database struct {
	// Type 类型
	Type        string `mapstructure:"type"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"prefix"`
}

var DatabaseSetting = &Database{}

// Setup 设置
func Setup() {
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
		panic(fmt.Errorf("读取配置异常: %w", err))
	}
	fmt.Println("初始化配置文件成功")
	viper.WatchConfig()

	// 将配置信息解析为实体
	err = viper.UnmarshalKey("settings", Setting)
	if err != nil {
		panic(fmt.Errorf("读取配置异常，解析失败: %w", err))
	}

	// 设置为全局变量，后续有其他配置则新增实体和变量即可
	ServerSetting = &Setting.Server
	DatabaseSetting = &Setting.Database

	// 设置初始值
	// 超时时间单位设置为秒
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}
