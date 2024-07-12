package request

type Category struct {
	Name string `json:"name" binding:"required,min=4,max=2048"`
	Type string `json:"type" binding:"required,min=1,max=255"`
	Slug string `json:"-"`
}
