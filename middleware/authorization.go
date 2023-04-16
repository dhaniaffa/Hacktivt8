package middleware

import (
	"myGram/database"
	"myGram/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func CommentAutorisasi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		itemId, err := strconv.Atoi(ctx.Param("commentId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})

			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))

		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(itemId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Data Not Found",
			})

			return
		}

		if Comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Unauthorized",
			})
		}

		ctx.Next()
	}
}

func PhotoAutorisasi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		itemId, err := strconv.Atoi(ctx.Param("photoId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})

			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))

		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(itemId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Data Not Found",
			})

			return
		}

		if Photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Unauthorized",
			})
		}

		ctx.Next()
	}
}

func SocialMediaAutorisasi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		itemId, err := strconv.Atoi(ctx.Param("sosmedId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})

			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))

		Sosmed := models.SocialMedia{}

		err = db.Select("user_id").First(&Sosmed, uint(itemId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Data Not Found",
			})

			return
		}

		if Sosmed.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Unauthorized",
			})
		}

		ctx.Next()
	}
}
