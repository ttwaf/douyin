package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义一个全局对象db
var db *gorm.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:douyin123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	//ref: https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭MySQL连接 Gorm中不需要关心数据库连接关闭问题
//func Close() {
//	_ = db.Close()
//}
