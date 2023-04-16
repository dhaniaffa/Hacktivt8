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

// GetAllComment godoc
// @Summary Get All Comment Endpoint
// @Schemes
// @Description Get All Comment By User
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comment [get]
func GetAllComment(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Item := []models.Comment{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Find(&Item)

	if err.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err":     "Not Found",
			"message": "Comments Not Found",
		})

		return
	}

	ctx.JSON(http.StatusFound, gin.H{
		"message": "Data Found",
		"count":   err.RowsAffected,
		"data":    Item,
	})
}

// GetOneComment godoc
// @Summary Get One Comment Endpoint
// @Schemes
// @Description Get One Comment By Post Photo & User
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Param commentId path uint true "ID of the comment"
// @Router /comment/{commentId} [get]
func GetOneComment(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	Item := models.Comment{}

	ItemID, _ := strconv.Atoi(ctx.Param("commentId"))

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

// CreateComment godoc
// @Summary Create Comment Endpoint
// @Schemes
// @Description Create Comment
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Param photoId path uint true "ID of the photo"
// @Router /comment/post/{photoId} [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(ctx)

	Item := models.Comment{}

	ItemPhoto := models.Photo{}

	PhotoId, _ := strconv.Atoi(ctx.Param("photoId"))

	errPhoto := db.First(&ItemPhoto, uint(PhotoId)).Error

	if errPhoto != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err":     errPhoto.Error(),
			"message": "Photos Not Found",
		})
	}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Item)
	} else {
		ctx.ShouldBind(&Item)
	}

	Item.UserID = userID

	Item.PhotoID = ItemPhoto.ID

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

// UpdateComment godoc
// @Summary Update Comment Endpoint
// @Schemes
// @Description Update Comment By ID
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Param commentId path uint true "ID of the comment"
// @Router /comment/{commentId} [put]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()

	userData := ctx.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(ctx)

	Item := models.Comment{}

	ItemId, _ := strconv.Atoi(ctx.Param("commentId"))

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Item)
	} else {
		ctx.ShouldBind(&Item)
	}

	Item.UserID = userID

	Item.ID = uint(ItemId)

	err := db.Model(&Item).Where("id = ?", ItemId).Updates(models.Comment{
		Message: Item.Message,
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

// DeleteComment godoc
// @Summary Delete Comment Endpoint
// @Schemes
// @Description Delete Comment By ID
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Param commentId path uint true "ID of the comment"
// @Router /comment/{commentId} [delete]
func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()

	_ = ctx.MustGet("userData").(jwt.MapClaims)

	Item := models.Comment{}

	ItemId, _ := strconv.Atoi(ctx.Param("commentId"))

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
