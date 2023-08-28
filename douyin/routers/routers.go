package routers

import (
	"douyin/controller"
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {


	r := gin.Default()//创建局部路由
	//创建路由组
	douyin := r.Group("/douyin")
	douyin.POST("/user/register",controller.SignupHandler)//注册业务的路由
	douyin.POST("/user/login",controller.LoginHandler)//登录业务的路由
	//中间件开发。登录之前可以用下面的功能
	//登录后会认证。token正确的话就可以访问这些功能，不正确的话就打回去
	douyin.Use(middlewares.JWTAuthMiddleware())//应用JWT中间件
	{

	}

	return r
}

