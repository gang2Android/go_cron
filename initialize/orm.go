package initialize

import (
	"cronProject/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GetDB() *gorm.DB {
	return MysqlDB()
}

func MysqlDB() *gorm.DB {
	dsn := global.Config.Mysql.Name + ":" + global.Config.Mysql.Pwd + "@tcp(" + global.Config.Mysql.Host + ":" +
		global.Config.Mysql.Port + ")/" + global.Config.Mysql.Db + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN: dsn,
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, Logger: newLogger})
	//db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Error)
	return db
}
