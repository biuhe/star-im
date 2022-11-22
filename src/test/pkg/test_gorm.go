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
	fmt.Println("实体", prod2)

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
