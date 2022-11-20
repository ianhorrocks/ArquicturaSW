package productController

import (
	dto "ArquicturaSW/dto"
	service "ArquicturaSW/services/product"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, productDto)
		return
	}
	
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(http.StatusBadRequest, productsDto)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}
