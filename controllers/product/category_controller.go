package controllers

import (
	service "ArquicturaSW/services/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCategoryById(c *gin.Context) {
	log.Debug("Category id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var categoryDto = service.CategoryService.GetCategoryById(id)

	c.JSON(http.StatusOK, categoryDto)
}

func GetCategories(c *gin.Context) {
	var categoriesDto = service.CategoryService.GetCategories()

	c.JSON(http.StatusOK, categoriesDto)
}