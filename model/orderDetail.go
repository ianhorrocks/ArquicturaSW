package model

type OrderDetail struct {
	OrderDetailID  int    `gorm:"primary_key"`
	PrecioUnitario float64 `gorm:"type:decimal(60,4);not null"`
	Cantidad       int     `gorm:"type:integer(255);not null"`
	Nombre         string  `gorm:"type:varchar(350);not null"`
	OrderID        int    `gorm:"type:int(150);not null"`
	ProductID      int    `gorm:"type:int(150);not null"`
}

type OrderDetails []OrderDetail
