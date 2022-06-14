package controllers

import (
	service "ArquicturaSW/services/search"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetProductsBySearchParam(c *gin.Context) {

	param := c.Param("param")

	productsDto, err := service.SearchService.GetProductsBySearchParam(param)

	if err != nil {
		c.JSON(http.StatusBadRequest, productsDto)
		return
	}

	c.JSON(http.StatusAccepted, productsDto)
}
