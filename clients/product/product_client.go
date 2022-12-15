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

func GetProductsByText(texto string) model.Products {
	var products model.Products

	Db.Where("description LIKE ? OR name LIKE ?", "%"+texto+"%", "%"+texto+"%").Find(&products)
	log.Debug("Products: ", products)

	return products
}

func GetProductsByCategory(idCategory int) model.Products {
	var products model.Products

	Db.Where("id_category = ?", idCategory).Find(&products)
	log.Debug("Products: ", products)

	return products
}

func GetProductsByNumber(cantidad int) model.Products {
	var products model.Products
	Db.Order("RAND()").Limit(cantidad).Find(&products)

	log.Debug("Products: ", products)

	return products

}