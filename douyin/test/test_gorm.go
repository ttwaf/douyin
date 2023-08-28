package main

import (
	"douyin/models"
	//"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//要用自己的
//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}
//1、导包，复制之后改数据库open
func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema.这表示没类会自己创建.实体类。新增了一个字段的话就把后面注释掉重新运行testgorm
	//db.AutoMigrate(&Product{})
	db.AutoMigrate(&models.RegisterForm{})

	// Create。创建
	//db.Create(&Product{Code: "D42", Price: 100})
	user := &models.RegisterForm{}
	user.UserName = "阿西吧"
	db.Create(user)

	// Read
	//var product Product
	//db.First(&product, 1) // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//fmt.Println(db.First(user, 1))

	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	db.Model(user).Update("PassWord", "1234")//密码修改为1234


	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)
}
