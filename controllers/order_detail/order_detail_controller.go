package orderDetailController

import (
	"ArquicturaSW/dto"
	service "ArquicturaSW/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)
// LOS CONTROLLER TIENE QUE TENER EL MISMO DATO DEL DTO EJ 
func GetOrderDetailByIdOrder(c *gin.Context) {
	log.Debug("OrderDetail id to load: " + c.Param("OrderId"))

	OrderId, _ := strconv.Atoi(c.Param("idOrder")) //se pasa el id de array a int
	var orderDetailResDto dto.OrderDetailsDto

	orderDetailResDto, err := service.OrderDetailService.GetOrderDetailByIdOrder(OrderId) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDetailResDto)
}
