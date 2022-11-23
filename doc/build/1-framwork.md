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

在根目录中新增 src 目录，以及 main、resource、test 三个下级目录，用于存放主要程序文件、资源及配置文件、测试文件。

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

通过浏览器访问：http://localhost:8080/ping

得到如下信息：

```json
{
  "message": "pong"
}
```

此时我们就已经完成了http框架的测试，官方 GitHub 文档有提供不同请求方式、参数绑定、文件上传等示例，可以参考学习。后续将会基于野火IM
的[PC前端](https://github.com/wildfirechat/vue-pc-chat) 以及 [Java后端](https://github.com/wildfirechat/app-server)
示例进行改造。









