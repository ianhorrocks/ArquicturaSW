package services

import (
	orderDetailCliente "ArquicturaSW/clients/order_detail"
	"ArquicturaSW/dto"
	"ArquicturaSW/model"
	e "ArquicturaSW/utils/errors"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	GetOrderDetailByIdOrder(idOrder int) (dto.OrderDetailsDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) GetOrderDetailByIdOrder(idOrder int) (dto.OrderDetailsDto, e.ApiError) {

	var ordersDetail model.OrderDetails = orderDetailCliente.GetOrderDetailByIdOrder(idOrder)
	var ordersDetailResDto dto.OrderDetailsDto

	if len(ordersDetail) == 0 {
		return ordersDetailResDto, e.NewBadRequestApiError("orderDetail not found")
	}
	for _, orderDetailRes := range ordersDetail {
		var orderDetailResDto dto.OrderDetailDto
		orderDetailResDto.Id = orderDetailRes.Id
		orderDetailResDto.Nombre = orderDetailRes.Nombre
		orderDetailResDto.Cantidad = orderDetailRes.Cantidad
		orderDetailResDto.PrecioUnitario = orderDetailRes.PrecioUnitario
		orderDetailResDto.Total= orderDetailRes.Total
		orderDetailResDto.IdProducto = orderDetailRes.ProductID
		orderDetailResDto.IdOrder = orderDetailRes.OrderID
		ordersDetailResDto = append(ordersDetailResDto, orderDetailResDto)
	}
	return ordersDetailResDto, nil
}