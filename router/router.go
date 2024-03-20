package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.Use(middlewares.Auth()).PUT("/", controllers.UpdateUser)
		userRouter.Use(middlewares.Auth()).DELETE("/", controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())
		photoRouter.POST("/", controllers.PostPhoto)
		photoRouter.GET("/", controllers.GetUserPhotos)
		photoRouter.GET("/all", controllers.GetAllPhotos)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthz(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthz(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Auth())
		commentRouter.POST("/", controllers.PostComment)
		commentRouter.GET("/", controllers.GetUserComments)
		commentRouter.GET("/photo/:photoId", controllers.GetPhotoComments)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthz(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthz(), controllers.DeleteComment)
	}

	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.Auth())
		sosmedRouter.POST("/", controllers.PostSocialMedia)
		sosmedRouter.GET("/", controllers.GetSocialMedias)
		sosmedRouter.PUT("/:socialMediaId", middlewares.SosmedAuthz(), controllers.UpdateSocialMedia)
		sosmedRouter.DELETE("/:socialMediaId", middlewares.SosmedAuthz(), controllers.DeleteSocialMedia)
	}

	return r
}
