package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int            `gorm:"primaryKey;autoIncrement"`
	TgId      int64          `gorm:"tg_id:bigint;not null"`
	Name      string         `gorm:"name:varchar(255);not null"`
	Email     string         `gorm:"email:varchar(255);not null"`
	Password  string         `gorm:"password:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at:timestamp;default:null"`
	Budget    []Budget
}
