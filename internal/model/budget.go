package model

import (
	"gorm.io/gorm"
	"time"
)

type Budget struct {
	Id        int            `gorm:"primaryKey;autoIncrement"`
	Title     string         `gorm:"size:4096;not null"`
	Type      string         `gorm:"type:varchar(255);not null"`
	Amount    float64        `gorm:"amount:bigint;not null"`
	CreatedAt time.Time      `gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at:timestamp;default:null"`
	UserID    uint           `gorm:"foreignkey:User REFERENCES users(id)"`
	User      User           `gorm:"association_foreignkey:ID"`
}
