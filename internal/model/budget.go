package model

import (
	"gorm.io/gorm"
	"time"
)

type Budget struct {
	Id        int            `gorm:"primaryKey;autoIncrement"`
	Type      string         `gorm:"type:varchar(255);not null"`
	Amount    float32        `gorm:"amount:bigint;not null"`
	CreatedAt time.Time      `gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at:timestamp;default:null"`
}
