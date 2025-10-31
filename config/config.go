package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string        //JWt 密钥
	JWTExpire  time.Duration //过期时间
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("警告 .env 文件没找到，使用环境变量")
	}

	expireHours, err := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "72"))

	if err != nil {
		expireHours = 72
	}

	AppConfig = Config{
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""), // 密码默认为空
		DBName:     getEnv("DB_NAME", "blog_go"),
		JWTSecret:  getEnv("JWT_SECRET", "default_secret"),
		JWTExpire:  time.Hour * time.Duration(expireHours),
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
