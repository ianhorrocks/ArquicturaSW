package model

type Category struct {
	Id   int    `gorm:"primaryKey"` //1 electronica
	Name string `gorm:"varchar(255);not null"`
}

type Categories []Category
