package model

type BudgetCategory struct {
	BudgetID   uint     `gorm:"primary_key"`
	CategoryID uint     `gorm:"primary_key"`
	Budget     Budget   `gorm:"foreignkey:BudgetID"`
	Category   Category `gorm:"foreignkey:CategoryID"`
}
