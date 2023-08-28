package logic

import (
	"douyin/dao/mysql"
	"douyin/models"
	"douyin/pkg/jwt"
	"douyin/pkg/snowflake"
	"errors"
)

func SignUp(userName, password string) (u *models.RegisterForm, error error) {
	//1、判断用户是否存在
	err := mysql.CheckUserExist(userName)
	if err != nil {
		return nil,err
	}
	//2、不存在的话生成uid
	userId, err := snowflake.GetID()
	if err != nil {
		return nil, errors.New("创建用户失败")
	}
	//生成atoken
	atoken, _, err := jwt.GenToken(userId,userName)
	if err != nil {
		return nil, err
	}
	//3、构造实例，保存到数据库
	u = &models.RegisterForm{
		UserID: userId,
		Atoken: atoken,
		UserName: userName,
		Password: password,
	}
	if err != mysql.InsertUser(u) {
		return nil, err
	}
	return
}

//func SignUp(p *models.RegisterForm) (error error)  {
//
//	//1、判断用户是否存在
//	err := mysql.CheckUserExist(p.UserName)
//	if err != nil {
//		//数据库查询出错
//		return err
//	}
//	//2、不存在的话就生成UID
//	userId, err := snowflake.GetID()
//	if err != nil {
//		return errors.New("创建用户失败")
//	}
//	//构造一个user实例
//	//u := models.User{
//	//	UserID:		userId,
//	//	UserName:	p.UserName,
//	//	Password:	p.Password,
//	//	//Name:		p.Name,
//	//}
//	u := models.User{
//		UserID:		userId,
//		UserName:	p.UserName,
//		Password:	p.Password,
//		//Name:		p.Name,
//	}
//	//保存到数据库。要加密的
//	return mysql.InsertUser(u)
//}

func Login(userName, password string) (u *models.RegisterForm, error error) {
	//放到user里
	u = &models.RegisterForm{
		UserName: userName,
		Password: password,
	}
	//获取uid
	userId, err := mysql.Login(u)
	//生成atoken
	atoken, _, err := jwt.GenToken(userId,userName)
	if err != nil {
		return nil, err
	}
	(*u).UserID = userId
	(*u).Atoken = atoken
	//返回
	return
}

//func Login(p *models.LoginForm) (atoken, rtoken string, error error )  {//在这第二次用到models.LoginForm。注册也是
//	//放到user里面
//	user := &models.User{
//		UserName: p.UserName,
//		Password: p.Password,
//	}
//	if err := mysql.Login(user); err != nil {
//		return "","", err
//	}
//	//生成jwt
//	return jwt.GenToken(user.UserID,user.UserName)
//}