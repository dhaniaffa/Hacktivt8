package controller

import (
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// @BasePath /api/v1

// UserRegister godoc
// @Summary User Register Endpoint
// @Schemes
// @Description User Register
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":  "Gagal Create User",
			"Validasi": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, "Register Berhasil")
	}
}

// UserLogin godoc
// @Summary User Login Endpoint
// @Schemes
// @Description User Login
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("username = ?", User.Username).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Username atau Password Salah!",
		})
		return
	}

	comparePwd := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePwd {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Username atau Password Salah!",
		})

		return
	}

	token := helpers.GenerateToken(User.ID, User.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
