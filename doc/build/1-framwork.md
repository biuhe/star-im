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

