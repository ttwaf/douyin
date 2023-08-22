package models

type RegisterForm struct {
	UserName string `json:"username" binding:"required"` //是json类型的，映射到小username，且要不为空。注意是反引号1左边
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"` //官方的接口里没有name，我给注释掉了
	//ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type DouyinUserRegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
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

type Users struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  int64  `json:"total_favorited,omitempty"`
	WorkCount       int64  `json:"work_count,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
}
