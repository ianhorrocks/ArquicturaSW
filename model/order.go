package model

import (
	"time"
)

type Order struct {
	OrderID int      `gorm:"primaryKey"`
	Amount  float64   `gorm:"type:decimal(10,2);not null"`
	Fecha   time.Time `gorm:"not null"`
	UserID  int      `gorm:"type:integer;not null"`
}

type Orders []Order
