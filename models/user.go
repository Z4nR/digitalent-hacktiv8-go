package models

import (
	"fmt"
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null" json:"username" valid:"required~Silahkan masukan nama pengguna anda"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" valid:"required~Silahkan masukan email yang ingin anda gunakan,email~Maaf format yang anda gunakan tidak sesuai"`
	Password     string        `gorm:"not null" json:"password" valid:"required~Silahkan masukan kata sandi anda,stringlength(8|20)~Kata sandi terdiri dari 8 sampai 20 karakter"`
	Age          int           `gorm:"not null" json:"age" valid:"required~Silahkan masukan usia anda,int~Maaf jenis data tidak sesuai"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 8 {
		return fmt.Errorf("maaf usia anda terlalu muda")
	}

	if _, err := govalidator.ValidateStruct(u); err != nil {
		return fmt.Errorf("%w", err)
	}

	u.Password = helpers.HashPass(u.Password)
	fmt.Println(u.Password)
	return nil
}
