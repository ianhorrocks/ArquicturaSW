package controllers

import (
	dto "ArquicturaSW/dto"
	service "ArquicturaSW/services/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //El guion bajo es porque puede dar un error (variable anonima)
	var userDto dto.UserDto

	userDto, err := service.UserService.GetUserById(id) //El service tiene un metodo Get userbyid. Tiene la posibilidad de un error o un userDTO. La capa service se va encargar de buscar ese usuario

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) { 
	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}