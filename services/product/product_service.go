package services

import (
	productCliente "ArquicturaSW/clients/product"
	"ArquicturaSW/dto"
	"ArquicturaSW/model"
	e "ArquicturaSW/utils/errors"
)

type productService struct{}

type productServiceInterface interface { //nombre del metodo, que .. todo lo que va a devolver mi servicio es en DTO (ej: product)
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if product.ProductID == 0 {
		return productDto, e.NewBadRequestApiError("Product not found")
	}

	productDto.Id = product.ProductID
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Price = product.Price
	productDto.Stock = product.Stock

	return productDto, nil
}
func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	if len(products) == 0 {
		return productsDto, e.NewBadRequestApiError("Products not found")
	}
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id = product.ProductID
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Picture = product.Picture
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.IdCategory = product.CategoryID

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}