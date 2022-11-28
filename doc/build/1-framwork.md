# 搭建步骤1-框架相关

## go modules 包管理工具

### 定义

依赖管理工具，可以理解为 maven / gradle 等工具

[官方文档及介绍](https://github.com/golang/go/wiki/Modules)

> 模块是相关Go包的集合。modules是源代码交换和版本控制的单元。go命令直接支持使用modules，包括记录和解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。



### 初始化

在项目根目录（star-im/）中执行

``` go
go mod init star-im
```

会在根目录生成一个 go.mod 的文件来进行包依赖的管理，其中会包含我们所需要的依赖及版本内容，此外某些依赖后面会有 indirect
字样，表示该依赖为传递依赖，也就是非直接依赖。

其他命令如：

```text
go mod <command>
download    download modules to local cache 	-- 将模块下载到本地缓存
edit        edit go.mod from tools or scripts -- 从工具或脚本编辑go.mod
graph       print module requirement graph 		-- 打印模块需求图
init        initialize new module in current directory 	-- 初始化当前目录中的新模块
tidy        add missing and remove unused modules 			-- 添加缺少的模块并删除未使用的模块
vendor      make vendored copy of dependencies 					-- 制作依赖项的供应商副本
verify      verify dependencies have expected content 	-- 验证依赖项是否具有预期内容
why         explain why packages or modules are needed 	-- 解释为什么需要包或模块
```

比较常用的是 `init`, `tidy`, `edit` ，当我们引入依赖包的之后，可以使用 `go mod tidy `
来命令来整理依赖模块。其他更多内容可参考：[go mod使用 | 全网最详细](https://zhuanlan.zhihu.com/p/482014524)

可以使用命令 `go list -m -u all `来检查可以升级的package

依赖包仓库地址：https://pkg.go.dev/（相当于maven的 https://mvnrepository.com/），搜索需要的依赖包可以访问此链接，里面也包含了依赖包的使用事项等。

目前 go web 似乎没有比较成型的 web 开发标准，因此我沿用了 Java 的习惯

在根目录中新增 src 目录，以及 main、resource、test 三个下级目录，用于存放主要程序文件、资源设置文件、测试文件。

``` go
star-im
  └── src
      ├── main
      ├── resource
      └── test
```



## viper 配置管理库

### 定义

Viper是一个完整的Go应用程序配置解决方案，可以用于读取 JSON、TOML、YAML、HCL、env file和Java properties 配置文件。

我们通常将一些配置信息，如数据库的访问路径，端口号等存放在配置文件中方便统一修改。

在 Java 中通常为 `application.yml` 或者 `applicatiton.properties` 文件，然后在 springboot 框架下使
用 `@ConfigurationProperties(prefix=”setting_name”) ` 或者 `@Value(“valueStr”)` 的形式来读取。

Viper 就是 go 用于做这一部分的工作类库

相关链接：

[GitHub](https://github.com/spf13/viper)

[PKG](https://pkg.go.dev/github.com/spf13/viper)



### 安装

在项目中打开命令行执行如下命令

```shell
go get github.com/spf13/viper
```



### 使用

在 `项目根目录/src/resource` 资源目录下新建一个 `app.yml` 配置文件，并写入以下配置项

``` yaml
settings:
  server:
    url: localhost
    port: 8081
```

注：我们现在约定 `settings ` 为配置项根节点，之后新增例如 `settings:database`
之类的节点，则是在settings下新增一个 `database `节点，而不是重复设置多一个 `settings`。 其他新增/修改项也遵循此说法。

如在settings下新增 `database `内容， 并修改 `server`下的端口号为9999，示例如下

错误示例为：

``` yaml
settings:
  server:
    url: localhost
    port: 8081
server:   
	port: 9999
settings:
	database:
		type: mysql
```

正确示例为：

``` yaml
settings:
  server:
    url: localhost
    port: 9999
	database:
		type: mysql
```

### 测试

在  `项目根目录/src/test` 目录下新建一个 `pkg` 目录，用于测试引入的第三方类库。在目录下新建 `test_viper.go` 测试文件

``` go
// main 方法必须使用 main 包
package main

// 引入依赖
import (
	"fmt"
	"github.com/spf13/viper"
)

// 主要执行的方法
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

```

执行程序后在控制台输出如下结果：

```go
初始化 app 配置成功
配置中的服务器地址及端口号为：localhost:8081
```

## gorm 对象关系映射框架

### 定义

gorm是全功能的ORM框架，包含了对不同数据库的增删改查操作，并支持事务、批量操作等等。

和`Java`的`hibernate`框架相似

相关链接：

[GitHub](https://github.com/go-gorm/gorm)

[GORM中文网](https://gorm.io/zh_CN/docs/index.html)

### 安装

```shell
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

### 测试

创建数据库的步骤忽略，我们约定数据库名称为star-im，用户名和密码均为root。

在  `项目根目录/src/test/pkg` 目录下新建 `test_gorm.go` 测试文件

``` go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TestProduct 定义一个实体
type TestProduct struct {
	// gorm.Model 提供了基础实体的定义，包含了id, CreatedAt, UpdatedAt, DeletedAt 字段
	gorm.Model
	// Name 商品名称
	Name string
	// Price 商品价格
	Price uint
}

// 测试 ORM 框架 —— 连接 MySQL https://github.com/go-gorm/gorm
func main() {
	// 连接信息，字符串中内容分别为：用户名:密码@连接方式(Host:Port)/数据库名?字符集&解析时间&默认时间
	// 更多参数详见：https://github.com/go-sql-driver/mysql#parameters
	dsn := "root:root@tcp(127.0.0.1:3306)/star-im?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接数据库，并设置基本的配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 如果连接有异常则打印
		fmt.Println("连接数据库失败：", err)
	}

	// 迁移 schema，如果数据库该表没有则创建表
	err = db.AutoMigrate(&TestProduct{})
	if err != nil {
		fmt.Println("创建数据库表异常：", err)
	}

	// Create 创建记录
	// 定义实体
	product := &TestProduct{Name: "奶茶", Price: 100}
	// 创建记录
	result := db.Create(product)
	// 创建成功后会返回插入数据的主键给实体赋值 ID
	fmt.Println("ID为：", product.ID)
	fmt.Println("如果有异常，则会输出：", result.Error)
	fmt.Println("返回插入记录的条数：", result.RowsAffected)

	// Find 查询
	prod := db.First(&product, "name = ?", "奶茶")
	fmt.Println("查询数:", prod.RowsAffected)

	// 查找后返回实体
	prod2 := TestProduct{}
	db.Where("name = ?", "奶茶").First(&prod2)
	fmt.Println("实体：", prod2)

	// Update - 修改
	// 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(TestProduct{Price: 200, Name: "蛋糕"})
	// 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Name": "蛋糕"})

	// Delete - 逻辑删除 product，会修改 deleted_at，标记为删除
	db.Delete(&product, 1)
}

```

执行程序后在控制台输出如下结果：

``` shell
ID为： 1
如果有异常，则会输出： <nil>
返回插入记录的条数： 1
查询数: 1
实体： {{1 2022-11-22 16:53:58.969 +0800 CST 2022-11-22 16:53:58.969 +0800 CST {0001-01-01 00:00:00 +0000 UTC false}} 奶茶 100}
```

其他更多操作请参考 [GORM中文网](https://gorm.io/zh_CN/docs/index.html)
，以及 [约束](https://gorm.io/zh_CN/docs/constraints.html)、[连接池](https://gorm.io/zh_CN/docs/generic_interface.html)
、[日志](https://gorm.io/zh_CN/docs/logger.html) 等配置可根据自身需求学习设置。我在后续编码过程中也会讲解并设置。

## gin http 框架

### 定义

Gin是用 Go 开发的一个HTTP web 微框架，类似 Martinier
的API，重点是小巧、易用、性能好很多，也因为 [httprouter](https://github.com/julienschmidt/httprouter) 的性能提高了40倍

相关链接：

[GitHub](https://github.com/gin-gonic/gin)

### 安装

``` shell
go get github.com/gin-gonic/gin
```

### 测试

在  `项目根目录/src/test/pkg` 目录下新建 `test_gin.go` 测试文件

``` go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

```

通过浏览器访问：http://localhost:8081/ping

得到如下信息：

```json
{
  "message": "pong"
}
```

此时我们就已经完成了http框架的测试，官方 GitHub
文档有提供不同请求方式、参数绑定、文件上传等示例，可以参考学习。后续将会基于野火IM的[PC前端](https://github.com/wildfirechat/vue-pc-chat)
以及 [Java后端](https://github.com/wildfirechat/app-server)示例进行改造。

## 搭建框架

通过上述步骤我们已经能基本了解框架所需要的基本内容，现在我们真正以实现登录功能为目标来打通整个流程。

和上述步骤一样，我们从配置项搭建开始，在  `项目根目录/src/main/` 目录下新建一个 `config`
目录，用于存放配置文件。在该目录下新建 `database`、`settings` 目录，并分别新建 `database.go` 和 `settings.go` 文件，用做初始化读取配置
—— viper、以及初始化数据库 —— gorm的操作。

### Viper 读取配置

`settings.go` 提供了读取配置并设置到全局实体提供给其他类使用，代码如下：

``` go
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

```

该类主要操作

1. 读取配置文件并解析为实体
2. 设置全局变量提供给其他类使用
3. 设置初始值

### Gorm 连接数据库

`database.go` 提供了初始化数据库连接的操作，代码如下：

```go
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

```

该类主要操作

1. 读取数据库连接配置
2. 初始化数据库连接
3. 定义了 慢 SQL 日志配置
4. 设置了数据库连接池配置
5. 返回全局变量DBS供其他类使用

### Init 加载配置

在  `项目根目录/src/main/config` 目录下新建 `init.go` 文件，用于初始化上面两个配置项。

`init.go` 代码如下：

``` go
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
  // 后续有其他配置项可以在下面添加……
}

```

该文件到时候放在main方法中执行即可

### main 程序入口

在  `项目根目录` 下新建一个 `main.go` 作为我们作为http程序主入口，参考 `gin` 章节初始化`gin`

`main.go` 代码如下：

```go
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

```

该类主要操作

1. 初始化配置
2. 初始化路由配置以及服务基础设置

`routers.Setup()` ，路由等信息单独放在另外一个目录 `routers`中来统一管理。

### router 路由配置

在  `项目根目录/src/main/` 目录下新建一个 `routers` 目录，并按照层级建立 `api/v1` 两个目录，用于存放路由接口。v1
表示接口版本号，方便后续迭代接口版本。

在 `项目根目录/src/main/routers/api` 目录下创建 `health.go` ，代码如下：

``` go
package api

import (
	"github.com/gin-gonic/gin"
	"star-im/src/main/common/app"
)

// Ping 接口连通性测试
func Ping(c *gin.Context) {
	// 直接返回成功结果
	c.JSON(http.StatusOK, gin.H{
			"msg": "成功",
		})
}
```

该类主要做连通性测试，因此直接返回json成功数据

在 `项目根目录/src/main/routers`  目录下创建 `router.go` 文件用于初始化路由配置，代码如下

``` go
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
  // 统一日志
	r.Use(gin.Logger())

	// 不需要鉴权
	r.GET("/ping", api.Ping)
  
	return r
}
```

该类主要操作

1. 初始化路由设置
2. 指定记录日志到文件
3. 指定具体的路由地址以及请求方式和响应函数

这时候我们启动根目录下的 `main`函数即可启动服务，通过浏览器访问：http://localhost:8081/ping 可以得到返回值

``` json
{"msg": "成功"}
```

### 统一返回值

我们规范约定返回值参数有利于我们对数据进行管理以及提升前后端开发的效率。

在  `项目根目录/src/main/` 目录下新建一个` common/app ` 层级目录，并在 `app `目录下分别建立 `code.go`、`msg.go`
、 `response.go` 用于存放 返回值、返回消息、统一返回值的实体对象

`code.go` 主要定义返回值常量，代码如下：

```go
package app

import "net/http"

const (
	SUCCESS = http.StatusOK
	ERROR   = -1
)
```

`msg.go` 主要定义返回值常量对应的消息内容，代码如下：

``` go
package app

// MessageMap 返回值常量对应的消息内容，消息集合：{消息码，消息内容}
var MessageMap = map[int]string{
	SUCCESS: "成功",
	ERROR:   "失败",
}

// GetMsg 根据代码获取返回信息
func GetMsg(code int) string {
	msg, ok := MessageMap[code]
	if ok {
		return msg
	}

	return MessageMap[ERROR]
}

```

`response.go` 主要定义返回值的对象，代码如下：

``` go
package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应对象
type Response struct {
	// 响应编码
	Code int `json:"code"`
	// 返回消息
	Msg string `json:"msg"`
	// 返回数据
	Data interface{} `json:"data"`
}

// Res 设置 gin.JSON 的内容
func Res(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}

// Success 返回成功结果
func Success(c *gin.Context, data interface{}) {
	Res(c, http.StatusOK, SUCCESS, data)
}

// Error 返回错误结果，异常结果放在统一异常处理 handler中
func Error(c *gin.Context, data interface{}) {
	Res(c, http.StatusOK, ERROR, data)
}

```

这时我们可以修改 `src/main/routers/api/health.go` 中返回的结果如下：

``` go
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

```

### 统一异常处理

我们需要统一处理系统的异常信息并让异常结果也显示为统一的返回结果对象，那么需要进行统一异常处理。

在  `项目根目录/src/main/` 目录下新建一个 `handler` 目录，并在目录下新建一个 `exception.go` 文件，用于处理异常信息

`exception.go` 代码如下：

``` go
package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"star-im/src/main/common/app"
)

// Recover 注意 Recover 要尽量放在router.User的第一个被加载
// 如不是的话，在recover前的中间件或路由，将不能被拦截到
// 程序的原理是：
// 1.请求进来，执行recover
// 2.程序异常，抛出panic
// 3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
func Recover(c *gin.Context) {
	defer func() {
		r := recover()
		if r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			c.JSON(http.StatusOK, app.Response{
				Code: app.ERROR,
				Msg:  ErrorToString(r),
				Data: nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()

	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// ErrorToString recover错误，转string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

```

该类主要捕获panic异常，并返回 json 信息给客户端

在 `src/main/routers/router.go` 文件中添加如下代码即可。

```go
// 统一异常处理
r.Use(handler.Recover)
```

参考：

[跟煎鱼学go](https://eddycjy.gitbook.io/golang/di-3-ke-gin/install)









