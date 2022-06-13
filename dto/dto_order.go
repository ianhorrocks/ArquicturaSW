package dto

type OrderDto struct {
	Id              uint    `json:"id"`
	UserId          uint    `json:"user_id"`
	AddressId       uint    `json:"address_id"`
	Amount          float64 `json:"amount"`
	OrderDetailsDto `json:"order_details"`
}

type OrdersDto []OrderDto
