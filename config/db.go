package config

import (
	"community/back/global"
	"log"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	// 连接数据库
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database, ", err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatal("Error connecting to database, ", err)
	}

	global.Db = db
	log.Println("Database connected successfully")
}
