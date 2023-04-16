package models

import (
	validator "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an comment
type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~message tidak boleh kosong!"`
	UserID  uint
	User    *User `gorm:"foreignKey:UserID"`
	PhotoID uint
	Photo   *Photo `gorm:"foreignKey:PhotoID"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := validator.ValidateStruct(comment)

	if errCreate != nil {
		err = errCreate

		return
	}

	err = nil

	return
}

func (comment *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := validator.ValidateStruct(comment)

	if errCreate != nil {
		err = errCreate

		return
	}

	err = nil

	return
}
