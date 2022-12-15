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
	var texto string = c.Param("texto")
	log.Debug("texto: ", texto)
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByText(texto)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByCategory(c *gin.Context) {
	log.Debug("Products categories id to load: " + c.Param("idCategory"))

	idCategory, _ := strconv.Atoi(c.Param("idCategory"))
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByCategory(idCategory)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}


func GetProductsByNumber(c *gin.Context) {
	cantidad, _ := strconv.Atoi(c.Param("cantidad"))
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByNumber(cantidad)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}