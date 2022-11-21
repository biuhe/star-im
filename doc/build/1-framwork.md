# 搭建步骤1-框架相关

## 初始化

### go modules 包管理工具

#### 定义

依赖管理工具，可以理解为 maven / gradle 等工具

[官方文档及介绍](https://github.com/golang/go/wiki/Modules)

> 模块是相关Go包的集合。modules是源代码交换和版本控制的单元。go命令直接支持使用modules，包括记录和解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。

#### 初始化

在项目根目录（star-im/）中执行

``` go
go mod init star-im
```

会在根目录生成一个 go.mod 的文件来进行包依赖的管理

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

目前 go web 似乎没有比较成型的 web 开发标准，因此我沿用了 Java 的习惯

在根目录中新增 src 目录，以及 main、resource、test 三个下级目录，用于存放主要程序文件、资源及配置文件、测试文件。

``` go
star-im
  └── src
      ├── main
      ├── resource
      └── test
```

### viper 配置管理库

#### 定义

Viper是一个完整的Go应用程序配置解决方案，可以用于读取 JSON、TOML、YAML、HCL、env file和Java properties 配置文件。

我们通常将一些配置信息，如数据库的访问路径，端口号等存放在配置文件中方便统一修改。

在 Java 中通常为 `application.yml` 或者 `applicatiton.properties` 文件，然后在 springboot 框架下使
用 `@ConfigurationProperties(prefix=”setting_name”) ` 或者 `@Value(“valueStr”)` 的形式来读取。

Viper 就是 go 用于做这一部分的工作类库

[GitHub](https://github.com/spf13/viper)

#### 使用

##### 下载

在项目中打开命令行执行如下命令

```shell
go get github.com/spf13/viper
```

##### 使用

在 `项目根目录/src/resource` 目录下新建一个 `app.yml` 文件，并写入以下配置项

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
	fmt.Println("配置中的服务器url、端口号为：", url+":"+port)
}

```







