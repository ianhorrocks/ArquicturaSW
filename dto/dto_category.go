package dto

type CategoryDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoriesDto []CategoryDto
