package middlewares

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.Debug().Select("id").First(&User, userID).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.ID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you aren't allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func PhotoAuthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Debug().Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you aren't allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func CommentAuthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Debug().Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you aren't allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func SosmedAuthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()
		sosmedId, err := strconv.Atoi(ctx.Param("socialMediaId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Sosmed := models.SocialMedia{}

		err = db.Debug().Select("user_id").First(&Sosmed, uint(sosmedId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Sosmed.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you aren't allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
