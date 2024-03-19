package controllers

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
		fmt.Println(User)
	} else {
		ctx.ShouldBind(&User)
		fmt.Println(User)
	}

	err := db.Create(&User).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
		fmt.Println(User)
	} else {
		ctx.ShouldBind(&User)
		fmt.Println(User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email or password",
		})
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	if token == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	userID := uint(userData["id"].(float64))

	var updatedUser models.User
	if contentType == appJSON {
		if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
	} else {
		if err := ctx.ShouldBind(&updatedUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
			return
		}
	}

	var existingUser models.User
	if err := db.First(&existingUser, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	existingUser.Email = updatedUser.Email
	existingUser.Username = updatedUser.Username

	if err := db.Save(&existingUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"age":       existingUser.Age,
		"email":     existingUser.Email,
		"username":  existingUser.Username,
		"id":        existingUser.ID,
		"update_at": existingUser.UpdatedAt,
	})
}
