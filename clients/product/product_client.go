package product

import (
	"ArquicturaSW/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProductById(id int) model.Product{
	var product model.Product

	Db.Where("id = ?", id).First(&product)
	log.Debug("Product: ", product)

	return product
}

func GetProducts() model.Products {
	var products model.Products

	Db.Find(&products)
	log.Debug("Products: ", products)

	return products
}
