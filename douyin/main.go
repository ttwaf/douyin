package main

import (
	"douyin/models"
	"douyin/pkg/snowflake"
	"douyin/routers"
	"fmt"
)

//localhost:8081
func main() {
	if err := models.InitDB(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	//defer mysql.Close() // 程序退出关闭数据库连接
	//gorm不需要

	//注册路由
	r := routers.SetupRouter() //注册全局的r路由。全放在里面

	err := r.Run(":8081")
	if err != nil {
		fmt.Printf("run server failed")
		return
	}
}
