package userController

import (
	"mvc-go/dto"
	service "mvc-go/services"
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

func GetUsers(c *gin.Context) { // No recibe ningun parametro
	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func OrderInsert(c *gin.Context) {
	//var orderDto dto.OrderDto
}

func UserInsert(c *gin.Context) {
	var userDto dto.UserDto 
	err := c.BindJSON(&userDto) // json que viene en el request (formulario) y lo convierte a userDto

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}
/*
debug
info
error
fatal
crash*/ 