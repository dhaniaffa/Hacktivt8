package models

import (
	validator "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for an photo
type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~title tidak boleh kosong!"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" valid:"required~caption tidak boleh kosong!"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo_url tidak boleh kosong!"`
	UserID   uint
	User     *User     `gorm:"foreignKey:UserID"`
	Comment  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"coments"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := validator.ValidateStruct(photo)

	if errCreate != nil {
		err = errCreate

		return
	}

	err = nil

	return
}

func (photo *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := validator.ValidateStruct(photo)

	if errCreate != nil {
		err = errCreate

		return
	}

	err = nil

	return
}
