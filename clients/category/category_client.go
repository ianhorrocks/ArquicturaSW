package category

import (
	model "ArquicturaSW/model"
	log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetCategories() model.Categories {
	var categories model.Categories

	Db.Find(&categories)

	log.Debug("Categories: ", categories)

	return categories
}
