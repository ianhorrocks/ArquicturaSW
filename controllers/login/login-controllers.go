package controllers

import (
	"net/http"

	dto "ArquicturaSW/dto"
	service "ArquicturaSW/services/login"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	cred := c.BindJSON(&loginDto)

	// Error Parsing json param
	if cred != nil {
		log.Error(cred.Error())
		c.JSON(http.StatusBadRequest, cred.Error())
		return
	}

	userDto, err := service.LoginService.Login(loginDto)
	// Error del Insert

	if err != nil {
		c.JSON(http.StatusBadRequest, userDto)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}
