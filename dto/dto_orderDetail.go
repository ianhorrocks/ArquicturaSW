package dto

type OrderDetailDto struct {
	Id             int     `json:"id"`
	IdOrder        int     `json:"id_order"`
	IdProducto     int     `json:"id_product"`
	Nombre         string  `json:"nombre"`
	PrecioUnitario float32 `json:"precio_unitario"`
	Cantidad       float32 `json:"cantidad"`
	Total          float32 `json:"total"`
}

type OrderDetailsDto []OrderDetailDto
