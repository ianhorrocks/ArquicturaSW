package model

type OrderDetail struct {
	Id             int     `gorm:"primaryKey"`
	PrecioUnitario float32 `gorm:"type:decimal(60,4);not null"`
	Cantidad       float32 `gorm:"type:decimal(60,4);not null"`
	Total          float32 `gorm:"type:decimal(60,4);not null"`
	Nombre         string  `gorm:"type:varchar(350);not null"`
	OrderID        int     `gorm:"type:int(150);not null"`
	ProductID      int     `gorm:"type:int(150);not null"`
}

type OrderDetails []OrderDetail
