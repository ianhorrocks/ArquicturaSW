package clients

import (
	model "ArquicturaSW/model"

	"errors"
	log "github.com/sirupsen/logrus"
	gorm "github.com/jinzhu/gorm"
)


var Db *gorm.DB

func GetProductsBySearchParam(ref string) (model.Products, error) {
	var products model.Products

	ref = "%" + ref + "%"

	InDescription := Db.Where("description LIKE ?", ref).Find(&products)

	err := InDescription.Error
	if err != nil {
		log.Println(err)
		return products, err
	}

	InName := Db.Where("name LIKE ?", ref).Find(&products)

	err = InName.Error

	if err != nil {
		log.Println(err)
		return products, err
	}

	if len(products) == 0 {
		err := errors.New("No encontramos el producto que estas buscando")
		return products, err
	}

	return products, nil
}
