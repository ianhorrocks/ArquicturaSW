package model

type Category struct {
	CategoryID int   `gorm:"primary_key"`  //1 electronica
	Name       string `gorm:"type:varchar(255);not null"`
}

type Categories []Category
