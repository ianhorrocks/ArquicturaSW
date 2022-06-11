package clients

import (
	"ArquicturaSW/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProducts() (model.Products, error) {
	var products model.Products

	err := Db.Find(&products).Error

	if err != nil {
		log.Println(err)
		return products, err
	}

	return products, err
}
