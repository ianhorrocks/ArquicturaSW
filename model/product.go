package model

type Product struct {
	ProductID   int    `gorm:"primary_Key"`
	Name        string  `gorm:"type:varchar(250);not null"`
	Price       float32 `gorm:"type:decimal(60,4);not null"`
	Picture		string	`gorm:"type:varchar(350);not null"`
	Stock       int     `gorm:"type:integer(250);not null"`
	Description string  `gorm:"type:varchar(250);not null"`
	CategoryID 	int    `gorm:"type:integer(150);not null"`
}

type Products []Product