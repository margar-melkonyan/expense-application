package seeder

import (
	"expense-application/internal/model"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type CategorySeeder struct {
	db *gorm.DB
}

func NewCategorySeeder(db *gorm.DB) *CategorySeeder {
	return &CategorySeeder{
		db: db,
	}
}

func (s *CategorySeeder) Seed() {
	s.createDefaultCategories()
}

func (s *CategorySeeder) createDefaultCategories() {
	categories := []map[string]string{
		{
			"Name": "Жилье",
			"Type": "expense",
		},
		{
			"Name": "Транспорт",
			"Type": "expense",
		},
		{
			"Name": "Питание",
			"Type": "expense",
		},
		{
			"Name": "Здоровье",
			"Type": "expense",
		},
		{
			"Name": "Страхование",
			"Type": "expense",
		},
		{
			"Name": "Образование",
			"Type": "expense",
		},
		{
			"Name": "Отдых и развлечения",
			"Type": "expense",
		},
		{
			"Name": "Подарки и благотворительность",
			"Type": "expense",
		},
		{
			"Name": "Другое (Расходы)",
			"Type": "expense",
		},
		{
			"Name": "Подработка",
			"Type": "income",
		},
		{
			"Name": "Заработная плата",
			"Type": "income",
		},
		{
			"Name": "Инвестиции",
			"Type": "income",
		},
		{
			"Name": "Другое (Доходы)",
			"Type": "income",
		},
	}

	for _, category := range categories {
		var ct = model.Category{
			Name: category["Name"],
			Slug: slug.Make(category["Name"]),
			Type: category["Type"],
		}

		if s.db.Table("categories").Where("slug = ?", ct.Slug).First(&ct).Error != nil {
			s.db.Create(&ct)
		}
	}
}
