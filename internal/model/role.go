package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id           uint           `json:"id" gorm:"primary_key;auto_increment"`
	Title        string         `json:"title" gorm:"varchar(100)"`
	DisplayTitle string         `json:"display_title" gorm:"varchar(100)"`
	Permissions  []byte         `json:"permissions" gorm:"jsonb"`
	CreatedAt    time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt    time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
}
