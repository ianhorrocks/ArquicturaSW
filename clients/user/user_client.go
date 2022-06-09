// clase DAO de user. La que se conecta con la Base de datos
package clients

import (
	model "ArquicturaSW/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB
//clases que insertan, buscan en el modelo de la bd

func GetUserById(id int) model.User { //GET
	var user model.User

	Db.Where("id_user = ?", id).First(&user) //Se transforma en un SQL y toma el primer registro. Todo eso lo hace la ORM. Db esta seteado en la linea 10. Ya vamos aprender como 
	log.Debug("User: ", user) //Objeto user con todos los datos seteados

	return user
}

func GetUsers() model.Users { //GET todo los usuarios de la base de datos
	var users model.Users
	Db.Find(&users) // Le pasa por referencia un puntero que acaba de crear

	log.Debug("Users: ", users)

	return users 
}

