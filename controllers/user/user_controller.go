package userController

import (
	"ArquicturaSW/dto"
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

func UserLogin(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenDto, er := service.UserService.LoginUser(loginDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, tokenDto)

}
