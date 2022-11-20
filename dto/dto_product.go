package dto

type ProductDto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture_url"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	IdCategory  int     `json:"id_categoty"`
}

type ProductsDto []ProductDto
