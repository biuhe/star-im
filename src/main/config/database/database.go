package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"star-im/src/main/config/settings"
	"time"
)

// DBS 定义全局变量，提供给其他方法调用
var DBS *gorm.DB

// Setup 初始化数据库连接
// https://gorm.io/zh_CN/
func Setup() {

	var err error
	//定义连接路径
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DatabaseSetting.User,
		settings.DatabaseSetting.Password,
		settings.DatabaseSetting.Host,
		settings.DatabaseSetting.Name)

	// 连接数据库，并设置基本的配置
	// 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			// 慢 SQL 阈值
			SlowThreshold: time.Second,
			// 日志级别
			LogLevel: logger.Silent,
			// 忽略ErrRecordNotFound（记录未找到）错误
			IgnoreRecordNotFoundError: true,
			// 彩色打印
			Colorful: true,
		},
	)
	DBS, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(fmt.Errorf("初始化数据库异常: %w", err))
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := DBS.DB()

	// 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
