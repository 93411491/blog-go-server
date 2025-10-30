package main

import (
	"blog-go/config"
	v1 "blog-go/internal/api/v1"
	"blog-go/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.Init()
	
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	{
		auth := apiV1.Group("/auth")
		{
			auth.POST("/register", v1.Register)
		}
	}

	r.Run(":8080")
}
