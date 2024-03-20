package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" valid:"required~Harap tambahkan nama anda"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" valid:"required~Harap tambahkan tautan sosmed anda"`
	UserID         uint   `json:"user_id"`
}

func (s *SocialMedia) ValidateComment(db *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		errCreate = err
		return
	}

	err = nil
	return
}
