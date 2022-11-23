package dto

type CategoryDto struct {
	IdCategory int    `json:"category_id"`
	Name       string `json:"name"`
}

type CategoriesDto []CategoryDto
