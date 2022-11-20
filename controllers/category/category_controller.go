package categoryController

import (
	"ArquicturaSW/dto"
	service "ArquicturaSW/services/category"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
)


func GetCategories(c *gin.Context) {
	var categoriesDto dto.CategoriesDto
	categoriesDto, err := service.CategoryService.GetCategories()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, categoriesDto)
}