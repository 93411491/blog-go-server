package database

import (
	"blog-go/config"
	"blog-go/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	cfg := config.AppConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("数据库连接失败:%v", err)
	}

	if err = DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("数据库迁移失败:%v", err)
	}

	log.Println("数据库连接成功")

}
