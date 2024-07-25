package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id                      uint           `json:"id" gorm:"primary_key;auto_increment" binding:"numeric,min=1"`
	Title                   string         `json:"title" gorm:"varchar(100)" binding:"required,alpha,min=1,max=255"`
	DisplayTitle            string         `json:"display_title" gorm:"varchar(100)" binding:"required,min=1,max=255"`
	Permissions             []byte         `json:"-" gorm:"jsonb"`
	PermissionsUnmarshalled []string       `json:"permissions" gorm:"-" binding:"required,dive,alpha,min=1,max=255"`
	CreatedAt               time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt               time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt               gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
}
