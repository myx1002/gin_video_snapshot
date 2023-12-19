package dao

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:941002@tcp(192.168.1.166:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			tmp := time.Now().Local().Format("2006-01-02 15:04:05")
			now, _ := time.ParseInLocation("2006-01-02 15:04:05", tmp, time.Local)
			return now
		},
	})
	return err
}
