package model

type BudgetCategory struct {
	BudgetID   int      `gorm:"primary_key"`
	CategoryID int      `gorm:"primary_key"`
	Budget     Budget   `gorm:"foreignkey:BudgetID"`
	Category   Category `gorm:"foreignkey:CategoryID"`
}
