package model

type Address struct { //clase Address
	StreetName string    `gorm:"type:varchar(350);not null"`
	StreetNumber int 	`gorm:"type:varchar(350);not null"`
	ZipCode	int			`gorm:"type:varchar(350);not null"`
}

type Addresses []Address