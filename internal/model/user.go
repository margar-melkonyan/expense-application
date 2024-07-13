package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	TgId      int64          `json:"tg_id" gorm:"tg_id:bigint;not null"`
	Name      string         `json:"-" gorm:"name:varchar(255);not null"`
	Email     string         `json:"-" gorm:"email:varchar(255);not null"`
	Password  string         `json:"-" gorm:"password:varchar(255);not null"`
	CreatedAt time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
	Budget    []Budget
}
