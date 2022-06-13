package model

type OrderDetail struct {
	OrderDetailID uint    `gorm:"primary_key"`
	Price         float64 `gorm:"type:decimal(10,2);not null"`
	Quantity      int     `gorm:"type:integer(255);not null"`
	OrderID       uint
	ProductID     uint
}

type OrderDetails []OrderDetail
