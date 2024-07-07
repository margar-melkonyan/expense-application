package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Slug      string         `gorm:"name:varchar(255);not null;unique"`
	CreatedAt time.Time      `gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at:timestamp;default:null"`
}
