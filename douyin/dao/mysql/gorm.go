package mysql

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	//自定义日志，打印sql语句，方便后期调试
	newLogger := logger.New(
		log.New(os.Stdout,"\r\n",log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,//sql阈值
			LogLevel: logger.Info,//级别
			Colorful: true,//彩色
		},
	)

	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		return err
	}
	return
	//fmt.Println("mysql inited.")
}