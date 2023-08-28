package mysql

import (
	"crypto/md5"
	"database/sql"
	"douyin/models"
	"encoding/hex"
	"errors"
	"fmt"
)

const secret = "douyin.vip"

func encryptPassword(data []byte) (result string) {//调用
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

func CheckUserExist(username string) (error error)  {
	//查username后求和，如果综合大于0，就说明已经存在了
	sqlStr := `select count(user_id) from users where username = ?`
	var count int
	//查是否出错
	if err := db.Get(&count, sqlStr, username); err != nil {//get是接收
		fmt.Println("查询出错")
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

//向数据库插入一个新用户
func InsertUser(user models.User) (error error) {
	//对密码进行加密md5
	user.Password = encryptPassword([]byte(user.Password))//密码拿出来生成新的密文密码之后再放回去
	//执行SQL语句入库
	sqlstr := `insert into users(user_id,username,password,name) values (?,?,?,?)`
	_, err := db.Exec(sqlstr,user.UserID,user.UserName,user.Password,user.Name)
	if err != nil {
		//fmt.Println("插入出错")
		return errors.New("插入出错")
	}
	return
}

func Login(user *models.User) (err error) {
	originPassword := user.Password
	sqlStr := `select user_id,username,password from users where username = ?`//用名字去查其它信息
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		//查询数据库出错
		return
	}
	if err == sql.ErrNoRows {
		//用户不存在
		return errors.New("用户不存在")
	}
	password := encryptPassword([]byte(originPassword))//再加密一遍
	if user.Password != password {
		return errors.New("密码错误")
	}
	return
}