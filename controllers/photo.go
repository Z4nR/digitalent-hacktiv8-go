package controllers

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostPhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	userID := uint(userData["id"].(float64))
	Photo := models.Photo{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Photo)
		fmt.Println(Photo)
	} else {
		ctx.ShouldBind(&Photo)
		fmt.Println(Photo)
	}

	Photo.UserID = userID

	if err := db.Debug().Create(&Photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"photo_url":  Photo.PhotoUrl,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"id":         Photo.ID,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func GetUserPhotos(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var (
		Photos []models.Photo
		User   models.User
		result []gin.H
	)

	if err := db.Debug().Where("user_id = ?", userID).Find(&Photos).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().First(&User, Photos[0].UserID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, photo := range Photos {
		result = append(result, gin.H{
			"id":         photo.ID,
			"created_at": photo.CreatedAt,
			"updated_at": photo.UpdatedAt,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoUrl,
			"User": gin.H{
				"username": User.Username,
				"email":    User.Email,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func GetAllPhotos(ctx *gin.Context) {
	db := database.GetDB()
	var (
		Photos []models.Photo
		result []gin.H
	)

	if err := db.Debug().Find(&Photos).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, photo := range Photos {
		var User models.User
		if err := db.Debug().Where("id = ?", photo.UserID).First(&User).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}

		result = append(result, gin.H{
			"id":         photo.ID,
			"created_at": photo.CreatedAt,
			"updated_at": photo.UpdatedAt,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoUrl,
			"User": gin.H{
				"username": User.Username,
				"email":    User.Email,
			},
		})
	}

	ctx.JSON(http.StatusAccepted, result)
}

func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Photo)
		fmt.Println(Photo)
	} else {
		ctx.ShouldBind(&Photo)
		fmt.Println(Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"updated_at": Photo.UpdatedAt,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
	})
}

func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	var (
		photo    models.Photo
		comments []models.Comment
	)

	if err := db.First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Data Tidak Ditemukan",
		})
		return
	}

	if err := db.Where("photo_id = ?", photoId).Delete(&comments).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Delete(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
