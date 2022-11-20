package model

type User struct { //clase usuario
	UserID   int   `gorm:"primary_Key"`
	Name     string `gorm:"type:varchar(250);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(250);not null;unique"` //Unique: No se puede repetir
	Password string `gorm:"type:varchar(250);not null"`
	Address  string `gorm:"type:varchar(350);not null"`
	Email    string `gorm:"type:varchar(250);not null"`

}

type Users []User //Coleccion de usuarios de la classe usuario
