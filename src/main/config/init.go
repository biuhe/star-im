package config

import (
	"star-im/src/main/config/database"
	"star-im/src/main/config/redis"
	"star-im/src/main/config/settings"
)

// Initial 初始化
func Initial() {
	// 初始化配置
	settings.Setup()
	// 初始化数据库连接
	database.Setup()
	// 初始化redis
	redis.Setup()
}
