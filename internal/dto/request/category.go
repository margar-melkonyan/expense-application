package request

type Category struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Slug string `json:"-"`
}
