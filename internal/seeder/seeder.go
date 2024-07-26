package seeder

import "gorm.io/gorm"

type User interface {
	Seed()
}

type Role interface {
	Seed()
}

type Category interface {
	Seed()
}

type Seeder struct {
	User
	Role
	Category
}

func (s Seeder) Seed() {
	s.Role.Seed()
	s.User.Seed()
	s.Category.Seed()
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{
		User:     NewUserSeeder(db),
		Role:     NewRoleSeeder(db),
		Category: NewCategorySeeder(db),
	}
}
