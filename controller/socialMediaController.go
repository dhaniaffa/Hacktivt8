package controller

import (
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

// @BasePath /api/v1

// GetAllSocialMedia godoc
// @Summary Get All Social Media Endpoint
// @Schemes
// @Description Get All Social Media User
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /social-media [get]
func GetAllSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Item := []models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Find(&Item)

	if err.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "Social Media Not Found",
		})

		return
	}

	ctx.JSON(http.StatusFound, gin.H{
		"message": "Data Found",
		"count":   err.RowsAffected,
		"data":    Item,
	})
}

// GetOneSocialMedia godoc
// @Summary Get All Social Media Endpoint
// @Schemes
// @Description Get One Social Media User
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Param sosmedId path uint true "ID of the social media"
// @Router /social-media/{sosmedId} [get]
func GetOneSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Item := models.SocialMedia{}

	ItemID, _ := strconv.Atoi(ctx.Param("sosmedId"))

	userID := uint(userData["id"].(float64))

	Item.UserID = userID

	err := db.First(&Item, uint(ItemID)).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "Data Not Found",
		})

		return
	}

	ctx.JSON(http.StatusFound, gin.H{
		"message": "Data Found",
		"data":    Item,
	})
}

// CreateSocialMedia godoc
// @Summary Create Social Media Endpoint
// @Schemes
// @Description Create Social Media
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /social-media [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(ctx)

	Item := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Item)
	} else {
		ctx.ShouldBind(&Item)
	}

	Item.UserID = userID

	err := db.Debug().Create(&Item).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create Success",
		"data":    Item,
	})
}

// UpdateSocialMedia godoc
// @Summary Update Social Media Endpoint
// @Schemes
// @Description Update Social Media By ID
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Param sosmedId path uint true "ID of the social media"
// @Router /social-media/{sosmedId} [put]
func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(ctx)

	Item := models.SocialMedia{}

	ItemId, _ := strconv.Atoi(ctx.Param("sosmedId"))

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Item)
	} else {
		ctx.ShouldBind(&Item)
	}

	Item.UserID = userID

	Item.ID = uint(ItemId)

	err := db.Model(&Item).Where("id = ?", ItemId).Updates(models.SocialMedia{
		Name:           Item.Name,
		SocialMediaURL: Item.SocialMediaURL,
	}).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
		"data":    Item,
	})
}

// DeleteSocialMedia godoc
// @Summary Delete Social Media Endpoint
// @Schemes
// @Description Delete Social Media By ID
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Param sosmedId path uint true "ID of the social media"
// @Router /social-media/{sosmedId} [delete]
func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()

	_ = ctx.MustGet("userData").(jwt.MapClaims)

	Item := models.SocialMedia{}

	ItemId, _ := strconv.Atoi(ctx.Param("sosmedId"))

	err := db.Model(&Item).Delete(&Item, ItemId).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}
