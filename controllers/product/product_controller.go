package productController

import (
	dto "ArquicturaSW/dto"
	service "ArquicturaSW/services"
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
		c.JSON(err.Status(), err)
		return
	}
	
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByText(c *gin.Context) {
	var text string = c.Param("texto")
	log.Debug("texto: ", text)
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByText(text)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByCategory(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("categoryId"))

	IdCategory, _ := strconv.Atoi(c.Param("IdCategory"))
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByCategory(IdCategory)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
