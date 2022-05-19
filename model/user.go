package model

type User struct { //clase usuario
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(350);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(150);not null;unique"` //Unique: No se puede repetir 
	Password string `gorm:"type:varchar(150);not null"`
}

type Users []User //Coleccion de usuarios de la classe usuario
