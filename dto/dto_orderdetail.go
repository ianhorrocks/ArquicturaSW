package dto


type OrderdetailDto struct {
	Amount float32 `json: "quantity"`
	ActualPrice float32 `json: "price"`
	Product ProductDto `json: "product"`
}

type OrderdetailsDto []OrderdetailDto



// producto, cantidad, precio del producto actual (se actualiza), subtotal(puede ir como no)
// relacion entre orden y detail. La orden tiene una coleccion de detalles
// puede ser que el detalle tenga una orden