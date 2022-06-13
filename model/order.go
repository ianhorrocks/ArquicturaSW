package model

import (

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserID       uint
	Amount       float64 `gorm:"type:decimal(10,2);not null"`
	User         User
	OrderDetails OrderDetails
}

type Orders []Order
