package services

import (
	productCliente "ArquicturaSW/clients/product"
	"ArquicturaSW/dto"
)

type productService struct{}

// ACA EL SERVICE SE COMUNICA CON EL MODEL(BUSINESS CLASSES) Y CONTROLLER(DTOs)
type productServiceInterface interface { //nombre del metodo, que .. todo lo que va a devolver mi servicio es en DTO (ej: product)
	GetProducts() (dto.ProductsDto, error)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProducts() (dto.ProductsDto, error) {

	products, err := productCliente.GetProducts()
	var productsDto dto.ProductsDto

	if err != nil {
		return productsDto, err
	}

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id = product.ProductID
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.Stock = product.Stock

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}