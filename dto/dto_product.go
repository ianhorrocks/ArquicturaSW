package dto

type ProductDto struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
}

type ProductsDto []ProductDto