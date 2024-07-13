package model

import (
	"gorm.io/gorm"
	"time"
)

type Budget struct {
	Id         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string         `json:"title" gorm:"size:4096;not null"`
	Type       string         `json:"type" gorm:"type:varchar(255);not null"`
	Amount     float64        `json:"amount" gorm:"amount:bigint;not null"`
	CreatedAt  time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt  time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
	UserID     uint           `json:"-" gorm:"foreignkey:User REFERENCES users(id)"`
	User       User           `json:"-" gorm:"association_foreignkey:ID"`
	Categories []*Category    `json:"-" gorm:"many2many:budget_categories;"`
}
