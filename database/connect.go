package database

import (
	"fiberTest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var  DB *gorm.DB
func Connect()  {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      false,         // 禁用彩色打印9
		},
		)
	database, err := gorm.Open(mysql.Open("root:123456@/go_main"), &gorm.Config{
		Logger: newLogger,
	})
	if err !=nil{
		log.Panic("无法连接数据库!")
	}
	DB = database
	DB.AutoMigrate(&models.User{},&models.Object{})
}