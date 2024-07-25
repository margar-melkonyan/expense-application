package seeder

import "gorm.io/gorm"

type User interface {
	CreateDefaultAdmin()
	Seed()
}

type Seeder struct {
	User
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{
		User: NewUserSeeder(db),
	}
}
