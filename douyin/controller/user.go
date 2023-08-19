package controller

import (
	"douyin/logic"
	"douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignupHandler(c *gin.Context)  {

	//1、从前端获取参数,2同时检验参数有效性
	var fo *models.RegisterForm//准备用fo接收
	if err := c.ShouldBindJSON(&fo); err != nil {

		c.JSON(http.StatusOK, gin.H{
			"msg":err.Error(),
		})

	}
	//2、注册用户的业务处理
	if err := logic.SignUp(fo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":"注册失败",
			"data":err.Error(),
		})
		return
	}
	//3、返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":"success",
	})
}

func LoginHandler(c *gin.Context)  {
	//1、获取请求参数及参数校验
	var u *models.LoginForm//拿到之后准备放到u里面
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误，直接返回相
		c.JSON(http.StatusOK, gin.H{
			"msg":err.Error(),
		})

	}
	//2、业务逻辑————登录
	atoken, _, err := logic.Login(u)
	if err != nil {

		if err.Error() == "用户不存在" {
			c.JSON(http.StatusOK, gin.H{
				"msg":"用户不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":"参数错误",
		})
		return
	}
	//3、返回响应
	ResponseSuccess(c, atoken)
}