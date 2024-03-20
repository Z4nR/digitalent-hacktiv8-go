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

func PostSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	userID := uint(userData["id"].(float64))
	Sosmed := models.SocialMedia{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Sosmed)
		fmt.Println(Sosmed)
	} else {
		ctx.ShouldBind(&Sosmed)
		fmt.Println(Sosmed)
	}

	Sosmed.UserID = userID

	if err := db.Debug().Create(&Sosmed).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"name":             Sosmed.Name,
		"id":               Sosmed.ID,
		"social_media_url": Sosmed.SocialMediaUrl,
		"user_id":          Sosmed.UserID,
		"created_at":       Sosmed.CreatedAt,
	})
}

func GetSocialMedias(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	var (
		Sosmeds []models.SocialMedia
		User    models.User
		result  []gin.H
	)

	if err := db.Debug().Where("user_id = ?", userID).Find(&Sosmeds).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().First(&User, Sosmeds[0].UserID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for _, sosmed := range Sosmeds {
		result = append(result, gin.H{
			"id":               sosmed.ID,
			"created_at":       sosmed.CreatedAt,
			"updated_at":       sosmed.UpdatedAt,
			"name":             sosmed.Name,
			"social_media_url": sosmed.SocialMediaUrl,
			"user_id":          sosmed.UserID,
			"User": gin.H{
				"username": User.Username,
				"email":    User.Email,
				"id":       User.ID,
			},
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"social_medias": result,
	})
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Sosmed := models.SocialMedia{}
	userID := uint(userData["id"].(float64))
	sosmedId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Sosmed)
		fmt.Println(Sosmed)
	} else {
		ctx.ShouldBind(&Sosmed)
		fmt.Println(Sosmed)
	}

	Sosmed.UserID = userID
	Sosmed.ID = uint(sosmedId)

	if err := db.Debug().Model(&Sosmed).Where("id = ?", sosmedId).Updates(models.SocialMedia{Name: Sosmed.Name, SocialMediaUrl: Sosmed.SocialMediaUrl}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":               Sosmed.ID,
		"updated_at":       Sosmed.UpdatedAt,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.SocialMediaUrl,
		"user_id":          Sosmed.UserID,
	})
}
