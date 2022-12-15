package services

import (
	orderCliente "ArquicturaSW/clients/order"
	orderDetailCliente "ArquicturaSW/clients/order_detail"
	productCliente "ArquicturaSW/clients/product"
	"ArquicturaSW/dto"
	"ArquicturaSW/model"
	e "ArquicturaSW/utils/errors"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
	GetOrdersByIdUser(token string) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {
	var order model.Order
	
	order.Fecha = time.Now()

	order.IdUser = orderDto.IdUsuario

	order = orderCliente.InsertOrder(order)

	orderDto.Id = order.Id

	var ordersDetail model.OrderDetails
	var montofinal float32
	for _, orderDetailDto := range orderDto.OrdersDetail {
		var orderDetail model.OrderDetail //

		orderDetail.ProductID = orderDetailDto.IdProducto

		var product model.Product = productCliente.GetProductById(orderDetail.ProductID)
		orderDetail.Nombre = product.Name
		orderDetail.PrecioUnitario = product.Price

		orderDetail.Cantidad = orderDetailDto.Cantidad
		orderDetail.Total = orderDetail.PrecioUnitario * orderDetail.Cantidad

		montofinal += orderDetail.Total

		orderDetail.OrderID = orderDto.Id

		ordersDetail = append(ordersDetail, orderDetail)
	}

	log.Debug("MONTO FINALLLL", montofinal)
	order.MontoFinal = orderCliente.UpdateMonto(montofinal, order.Id)

	log.Debug("Y ACAAA???", order.MontoFinal)

	ordersDetail = orderDetailCliente.InsertOrdersDetail(ordersDetail)

	var orderResponseDto dto.OrderDto

	orderResponseDto.Fecha = order.Fecha
	orderResponseDto.Id = order.Id
	orderResponseDto.IdUsuario = order.IdUser
	orderResponseDto.MontoFinal = order.MontoFinal

	log.Debug("order responde monto final",orderResponseDto.MontoFinal)
	

	for _, orderDetail := range ordersDetail {
		var orderDetailDto dto.OrderDetailDto

		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Nombre = orderDetail.Nombre
		orderDetailDto.IdOrder= orderDetail.OrderID
		orderDetailDto.PrecioUnitario = orderDetail.PrecioUnitario
		orderDetailDto.Total = orderDetail.Total
		orderDetailDto.Cantidad = orderDetail.Cantidad
		orderDetailDto.IdProducto = orderDetail.ProductID

		orderResponseDto.OrdersDetail = append(orderResponseDto.OrdersDetail, orderDetailDto)
	}
	return orderResponseDto, nil
}

func (s *orderService) GetOrdersByIdUser(token string) (dto.OrdersDto, e.ApiError) {
	var idUser float64
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, e.NewUnauthorizedApiError("No autorization")
		}
		return nil, e.NewUnauthorizedApiError("No autorization")
	}

	if !tkn.Valid {
		return nil, e.NewUnauthorizedApiError("No authorization")
	}
	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {

		idUser = (claims["id_user"].(float64))

	} else {
		return nil, e.NewUnauthorizedApiError("No authorization")
	}
	var IdUserX int = int(idUser)
	var orders model.Orders = orderCliente.GetOrdersByUser(IdUserX)
	var ordersDto dto.OrdersDto

	if len(orders) == 0 {
		return ordersDto, e.NewBadRequestApiError("order not found")
	}

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id
		orderDto.Fecha = order.Fecha
		orderDto.MontoFinal = order.MontoFinal

		var ordersDetail model.OrderDetails = orderDetailCliente.GetOrderDetailByIdOrder(order.Id)
		for _, orderDetail := range ordersDetail {
			var orderDetailDto dto.OrderDetailDto

			orderDetailDto.Id = orderDetail.Id
			orderDetailDto.Nombre = orderDetail.Nombre
			orderDetailDto.IdOrder = orderDetail.OrderID
			orderDetailDto.PrecioUnitario = orderDetail.PrecioUnitario
			orderDetailDto.Cantidad = orderDetail.Cantidad
			orderDetailDto.Total = orderDetail.Total
			orderDetailDto.IdProducto = orderDetail.ProductID

			orderDto.OrdersDetail = append(orderDto.OrdersDetail, orderDetailDto)
		}
		ordersDto = append(ordersDto, orderDto)
	}
	return ordersDto, nil

}
