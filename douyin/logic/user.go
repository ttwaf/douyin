package logic

import (
	"douyin/dao/mysql"
	"douyin/models"
	"douyin/pkg/jwt"
	"douyin/pkg/snowflake"
	"errors"
)

func SignUp(p *models.RegisterForm) (error error)  {

	//1、判断用户是否存在
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		//数据库查询出错
		return err
	}
	//2、不存在的话就生成UID
	userId, err := snowflake.GetID()
	if err != nil {
		return errors.New("创建用户失败")
	}
	//构造一个user实例
	u := models.User{
		UserID:		userId,
		UserName:	p.UserName,
		Password:	p.Password,
		Name:		p.Name,
	}
	//保存到数据库。要加密的
	return mysql.InsertUser(u)
}

func Login(p *models.LoginForm) (atoken, rtoken string, error error )  {//在这第二次用到models.LoginForm。注册也是
	//放到user里面
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "","", err
	}
	//生成jwt
	return jwt.GenToken(user.UserID,user.UserName)
}