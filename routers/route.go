package routers

import (
	"myGram/controller"
	_ "myGram/docs"
	"myGram/middleware"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Documentation myGram Final Project Hactivt8
// @version         1.0
// @description     Dokumentasi Tugas FInal Projek
// @termsOfService  http://swagger.io/terms/

// @contact.name   Dhaniaffa Adhimastama Mahadika
// @contact.url    https://github.com/dhaniaffa
// @contact.email  dhaniaffa@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// securitydefinitions.oauth2.accessCode OAuth2AccessCode

func ServeApp() *gin.Engine {
	route := gin.Default()

	userRoute := route.Group("/user")
	{
		userRoute.POST("/register", controller.UserRegister)

		userRoute.POST("/login", controller.UserLogin)
	}

	photoRoute := route.Group("/photo")
	{
		photoRoute.Use(middleware.Authentication())

		photoRoute.GET("/", controller.GetAllPhoto)

		photoRoute.POST("/", controller.CreatePhoto)

		photoRoute.GET("/:photoId", controller.GetOnePhoto)

		photoRoute.PUT("/:photoId", middleware.PhotoAutorisasi(), controller.UpdatePhoto)

		photoRoute.DELETE("/:photoId", middleware.PhotoAutorisasi(), controller.DeletePhoto)
	}

	commentRoute := route.Group("/comment")
	{
		commentRoute.Use(middleware.Authentication())

		commentRoute.GET("/", controller.GetAllComment)

		commentRoute.POST("/post/:photoId", controller.CreateComment)

		commentRoute.GET("/:commentId", controller.GetOneComment)

		commentRoute.PUT("/:commentId", middleware.CommentAutorisasi(), controller.UpdateComment)

		commentRoute.DELETE("/:commentId", middleware.CommentAutorisasi(), controller.DeleteComment)
	}

	socialMediaRoute := route.Group("/social-media")
	{
		socialMediaRoute.Use(middleware.Authentication())

		socialMediaRoute.POST("/", controller.CreateSocialMedia)

		socialMediaRoute.GET("/", controller.GetAllSocialMedia)

		socialMediaRoute.GET("/:sosmedId", controller.GetOneSocialMedia)

		socialMediaRoute.PUT("/:sosmedId", middleware.SocialMediaAutorisasi(), controller.UpdateSocialMedia)

		socialMediaRoute.DELETE("/:sosmedId", middleware.SocialMediaAutorisasi(), controller.DeleteSocialMedia)
	}

	route.GET("/docs/*any", ginSwagger.WrapHandler(files.Handler))

	return route
}
