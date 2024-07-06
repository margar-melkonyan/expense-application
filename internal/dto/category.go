package dto

type Category struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
