package v1

import (
	"blog-go/internal/model"
	"blog-go/pkg/database"
	"blog-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数无效：" + err.Error(),
		})
		return
	}
	var existingUser model.User

	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	newUser := model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "注册成功",
		"user": gin.H {
			"id": newUser.ID,
			"username":newUser.Username,
			"email":newUser.Email,
		},
	})

}
