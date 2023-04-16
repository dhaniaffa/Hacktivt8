package models

import (
	"myGram/helpers"

	validator "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an user
type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username tidak boleh kosong!"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email tidak boleh kosong!,email~Format E-mail salah!"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~Password tidak boleh kosong!,minstringlength(6)~Password minimal 6 karakter!"`
	Age         int           `gorm:"not null" json:"age" form:"age" valid:"required~Umur tidak boleh kosong!,numeric~Tidak dapat menerima karaker selain angka!,range(9|150)~Umur minimal diatas 8 tahun!"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"coments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := validator.ValidateStruct(user)

	if errCreate != nil {
		err = errCreate

		return
	}

	user.Password = helpers.HashPass(user.Password)

	err = nil

	return
}
