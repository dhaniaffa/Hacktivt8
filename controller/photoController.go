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

// GetAllPhoto godoc
// @Summary Get All Photo Endpoint
// @Schemes
// @Description Get All Photo By User
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photo [get]
func GetAllPhoto(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Item := []models.Photo{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Find(&Item)

	if err.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "Photos Not Found",
		})

		return
	}

	ctx.JSON(http.StatusFound, gin.H{
		"message": "Data Found",
		"count":   err.RowsAffected,
		"data":    Item,
	})
}

// GetOnePhoto godoc
// @Summary Get One Photo Endpoint
// @Schemes
// @Description Get One Photo By User
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Param photoId path uint true "ID of the photo"
// @Router /social-media/{photoId} [get]
func GetOnePhoto(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Item := models.Photo{}

	ItemID, _ := strconv.Atoi(ctx.Param("photoId"))

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

// CreatePhoto godoc
// @Summary Create Photo Endpoint
// @Schemes
// @Description Create Photo
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photo [post]
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(ctx)

	Item := models.Photo{}

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

// UpdatePhoto godoc
// @Summary Update Photo Endpoint
// @Schemes
// @Description Update Photo By ID
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Param photoId path uint true "ID of the photo"
// @Router /social-media/{photoId} [put]
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(ctx)

	Item := models.Photo{}

	ItemId, _ := strconv.Atoi(ctx.Param("photoId"))

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Item)
	} else {
		ctx.ShouldBind(&Item)
	}

	Item.UserID = userID

	Item.ID = uint(ItemId)

	err := db.Model(&Item).Where("id = ?", ItemId).Updates(models.Photo{
		Title:    Item.Title,
		PhotoURL: Item.PhotoURL,
		Caption:  Item.Caption,
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

// DeletePhoto godoc
// @Summary Delete Photo Endpoint
// @Schemes
// @Description Delete Photo By ID
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Param photoId path uint true "ID of the photo"
// @Router /social-media/{photoId} [delete]
func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()

	_ = ctx.MustGet("userData").(jwt.MapClaims)

	Item := models.Photo{}

	ItemId, _ := strconv.Atoi(ctx.Param("photoId"))

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
