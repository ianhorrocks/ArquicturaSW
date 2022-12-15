package dto

type CategoryDto struct {
	IdCategory int    `json:"id_category"`
	Name       string `json:"name"`
}

type CategoriesDto []CategoryDto
