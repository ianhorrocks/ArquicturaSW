package dto

type OrderDto struct {
	OrderDate string `json: "order_date"`
	Total float32 `json: "total"`
	User UserDto `json: "user"`
	OrderdetailDto OrderdetailDto `json: "detail"`
}

type OrdersDto []OrderDto


// usuario, total de orden int, status

//agregar un order controllers
//agregar un post en mapping