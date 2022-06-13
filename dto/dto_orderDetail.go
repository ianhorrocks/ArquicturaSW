package dto

type OrderDetailDto struct {
	Id        uint    `json:"id"`
	OrderId   uint    `json:"order_id"`
	ProductId uint    `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type OrderDetailsDto []OrderDetailDto
