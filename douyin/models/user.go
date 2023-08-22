package models

type RegisterForm struct {
	UserName string `json:"username" binding:"required"` //是json类型的，映射到小username，且要不为空。注意是反引号1左边
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"` //官方的接口里没有name，我给注释掉了
	//ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// 跟数据库绑定的结构体实例
type User struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"username"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name" db:"name"`
}

type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binging:"required"`
}
