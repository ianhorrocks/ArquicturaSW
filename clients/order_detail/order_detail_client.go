package orderDetail

import (
	"ArquicturaSW/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func InsertOrderDetail(orderDetail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&orderDetail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("OrderDetail ID ", orderDetail.Id, "created")
	result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.ProductID).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))

	if result1.Error != nil {
		log.Error("Pdto no encontrado")
	}

	return orderDetail
}

//Search order detail from an specific order
func GetOrderDetailByIdOrder(idOrder int) model.OrderDetails {
	var ordersDetail model.OrderDetails

	Db.Where("order_id = ?", idOrder).Find(&ordersDetail) // Fijarse si el id_order = ? corresponde al que esta en el dto o el modelo
	log.Debug("OrderDetail: ", ordersDetail)

	return ordersDetail
}

func InsertOrdersDetail(ordersDetail model.OrderDetails) model.OrderDetails {

	for _, orderDetail := range ordersDetail {

		log.Debug("ORDER DETAIL TO INSERT: ", orderDetail)

		result := Db.Create(&orderDetail)

		if result.Error != nil {
			log.Error("Error al crear")
		}
		log.Debug("OrderDetail Created: ", orderDetail.Id)
		result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.ProductID).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))

		if result1.Error != nil {
			log.Error("Pdto no encontrado")
		}
	}

	return ordersDetail
}

