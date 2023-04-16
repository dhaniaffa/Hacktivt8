package models

import (
	validator "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents the model for an socialmedia
type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~name tidak boleh kosong!"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social_media_url tidak boleh kosong!"`
	UserID         uint
	User           *User `gorm:"foreignKey:UserID"`
}

func (sosmed *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := validator.ValidateStruct(sosmed)

	if errCreate != nil {
		err = errCreate

		return
	}

	err = nil

	return
}

func (sosmed *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := validator.ValidateStruct(sosmed)

	if errCreate != nil {
		err = errCreate

		return
	}

	err = nil

	return
}
