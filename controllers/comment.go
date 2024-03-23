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

func PostComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	userID := uint(userData["id"].(float64))
	Comment := models.Comment{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Comment)
		fmt.Println(Comment)
	} else {
		ctx.ShouldBind(&Comment)
		fmt.Println(Comment)
	}

	Comment.UserID = userID

	if err := db.Debug().Create(&Comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetUserComments(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var (
		Comments []models.Comment
		User     models.User
		Photo    models.Photo
		result   []gin.H
	)

	if err := db.Debug().Where("user_id = ?", userID).Find(&Comments).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().First(&User, Comments[0].UserID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().First(&Photo, Comments[0].PhotoID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, comment := range Comments {
		result = append(result, gin.H{
			"id":         comment.ID,
			"created_at": comment.CreatedAt,
			"updated_at": comment.UpdatedAt,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"User": gin.H{
				"username": User.Username,
				"email":    User.Email,
				"id":       User.ID,
			},
			"Photo": gin.H{
				"id":        Photo.ID,
				"title":     Photo.Title,
				"caption":   Photo.Caption,
				"photo_url": Photo.PhotoUrl,
				"user_id":   Photo.UserID,
			},
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func GetPhotoComments(ctx *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	var (
		photo    models.Photo
		comments []models.Comment
		result   gin.H
	)

	if err := db.First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Where("photo_id = ?", photoId).Find(&comments).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	photo.Comments = comments

	result = gin.H{
		"photo_comments": photo,
	}

	ctx.JSON(http.StatusOK, result)
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Comment)
		fmt.Println(Comment)
	} else {
		ctx.ShouldBind(&Comment)
		fmt.Println(Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	if err := db.Debug().Model(&Comment).Where("id = ?", commentId).Update("message", &Comment.Message).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"updated_at": Comment.UpdatedAt,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
	})
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	if err := db.Debug().First(&Comment, commentId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().Delete(&Comment, commentId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
