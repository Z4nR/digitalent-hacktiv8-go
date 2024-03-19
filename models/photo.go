package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" valid:"required~Harap masukan judul foto"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" valid:"required~Harap masukan tautan gambar"`
	UserID   uint
}

func (p *Photo) ValidatePhoto(db *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		errCreate = err
		return
	}

	err = nil
	return
}
