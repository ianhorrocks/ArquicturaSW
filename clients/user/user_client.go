// clase DAO de user. La que se conecta con la Base de datos
package product

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB
//clases que insertan, buscan en el modelo de la bd

func GetUserById(id int) model.User { //GET
	var user model.User

	Db.Where("id = ?", id).First(&user) //Se transforma en un SQL y toma el primer registro. Todo eso lo hace la ORM. Db esta seteado en la linea 10. Ya vamos aprender como 
	log.Debug("User: ", user) //Objeto user con todos los datos seteados

	return user
}

func GetUsers() model.Users { //GET todo los usuarios de la base de datos
	var users model.Users
	Db.Find(&users) // Le pasa por referencia un puntero que acaba de crear

	log.Debug("Users: ", users)

	return users 
}

func InsertUser(user model.User) model.User { //POST, recibe como parametro el usuario que quiero insertar
	result := Db.Create(&user) // gorm traduce para crear e insertar un usuario

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}
