package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt *time.Time `gorm:"CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"CURRENT_TIMESTAMP"`
}
