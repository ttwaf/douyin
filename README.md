# douyin

数据库连接改了gorm

https://gorm.io/zh_CN/docs/sql_builder.html

dB.exec()
```go
var db *gorm.DB
func InitDB() (err error)
//gorm 初始化db
```

不建议使用dao目录

userID 在/pkg/jwt claim 中