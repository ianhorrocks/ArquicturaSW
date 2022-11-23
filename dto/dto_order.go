package dto

import "time"

type OrderDto struct {
	Id           int             `json:"id_order"`
	IdUsuario    int             `json:"id_user"`
	MontoFinal   float32         `json:"monto_final"`
	Fecha        time.Time       `json:"fecha"`
	OrdersDetail OrderDetailsDto `json:"OrdersDetalle"`
}

type OrdersDto []OrderDto
