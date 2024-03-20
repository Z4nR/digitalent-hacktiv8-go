package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `gorm:"CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"CURRENT_TIMESTAMP" json:"updated_at"`
}
