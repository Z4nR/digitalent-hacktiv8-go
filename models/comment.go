package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" valid:"required~Komentar Tidak Boleh Kosong"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}

func (c *Comment) ValidateComment(db *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		errCreate = err
		return
	}

	err = nil
	return
}
