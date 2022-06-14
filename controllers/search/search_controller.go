package controllers

import (
	service "ArquicturaSW/services/search"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetProductsBySearchParam(c *gin.Context) {

	ref := c.Param("ref")

	productsDto, err := service.SearchService.GetProductsBySearchParam(ref)

	if err != nil {
		c.JSON(http.StatusBadRequest, productsDto)
		return
	}

	c.JSON(http.StatusAccepted, productsDto)
}
