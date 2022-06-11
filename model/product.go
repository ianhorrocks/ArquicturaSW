package model

type Product struct {
	ProductID   uint    `gorm:"primary_Key"`
	Name        string  `gorm:"type:varchar(250);not null"`
	Description string  `gorm:"type:varchar(250);not null"`
	Price       float32 `gorm:"type:decimal(10,2);not null"`
	Stock       int     `gorm:"type:integer(250);not null"`
}

type Products []Product