package auth

import (
	"blog-go/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint, username string) (string, error) {
	claim := MyClaims{
		UserId:   userId,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.AppConfig.JWTExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "blog-go",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
