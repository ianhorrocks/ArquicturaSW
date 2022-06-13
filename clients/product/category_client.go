package clients

import (
	model "ArquicturaSW/model"
	log "github.com/sirupsen/logrus"
)

func GetCategoryById(id int) model.Category {
	var category model.Category

	Db.Where("category_id = ?", id).First(&category)
	log.Debug("Category: ", category)

	return category
}

func GetCategories() model.Categories {
	var categories model.Categories

	Db.Find(&categories)

	log.Debug("Products: ", categories)

	return categories
}
