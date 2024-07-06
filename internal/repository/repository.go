package repository

type Categories interface {
}

type Repository struct {
	Categories
}

func NewRepository() *Repository {
	return &Repository{}
}
