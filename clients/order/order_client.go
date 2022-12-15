package order

import (
	"ArquicturaSW/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Order Created: ", order.Id)
	return order
}

// Search order by idUser
func GetOrdersByUser(idUser int) model.Orders {
	var orders model.Orders

	log.Debug("Getting idUser: ", idUser)
	Db.Where("id_user = ?", idUser).Find(&orders)
	log.Debug("And orders: ", orders)

	return orders
}

// falta update monto final

func UpdateMonto(monto float32, idOrder int) float32 {
	result := Db.Model(&model.Order{}). Where("id = ?", idOrder).Update("monto_final", monto)

	log.Debug("ESTE ES EL MONTO DEL UPDATE", monto)
	
	if result.Error != nil{
		log.Error("Order no encontrada")
	}
	return monto
}