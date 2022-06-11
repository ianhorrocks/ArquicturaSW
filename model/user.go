package model

type User struct { //clase usuario
	UserID   uint   `gorm:"primary_Key;AUTO_INCREMENT"`
	Name     string `gorm:"type:varchar(250);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(250);not null;unique"` //Unique: No se puede repetir
	Email    string `gorm:"type:varchar(250);not null"`
	Password string `gorm:"type:varchar(250);not null"`
}

type Users []User //Coleccion de usuarios de la classe usuario
