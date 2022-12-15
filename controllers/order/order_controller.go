package orderController

import (
	"ArquicturaSW/dto"
	service "ArquicturaSW/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//Insertan orden

func OrderInsert(c *gin.Context) {
	var orderDto dto.OrderDto
	err := c.BindJSON(&orderDto)

	log.Debug("POSTING ORDER", orderDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDto)
}

//Buscar orden por token de usuarioX

func GetOrdersByIdUser(c *gin.Context) {

	var ordersDto dto.OrdersDto
	token := c.Param("idUser")
	ordersDto, err := service.OrderService.GetOrdersByIdUser(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersDto)
}