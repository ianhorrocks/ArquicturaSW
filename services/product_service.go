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
	GetProductsByText(texto string ) (dto.ProductsDto, e.ApiError)
	GetProductsByCategory(categoryId int) (dto.ProductsDto, e.ApiError)
	GetProductsByNumber(cantidad int) (dto.ProductsDto, e.ApiError)
	//get random products
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

// Get an specific product
func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("Product not found")
	}

	productDto.Id = product.Id
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Price = product.Price
	productDto.Stock = product.Stock
	productDto.IdCategory = product.IdCategory
	productDto.Picture = product.Picture

	return productDto, nil
}

// Display all products
func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	if len(products) == 0 {
		return productsDto, e.NewBadRequestApiError("Products not found")
	}
	for _, product := range products {
		var productDto dto.ProductDto

		productDto.Id = product.Id
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Picture = product.Picture
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

// Search product by text filter

func (s *productService) GetProductsByText(texto string) (dto.ProductsDto, e.ApiError) {
	var products model.Products = productCliente.GetProductsByText(texto)
	var productsDto dto.ProductsDto

	if len(products) == 0 {
		return productsDto, e.NewBadRequestApiError("Products not found")
	}

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id = product.Id
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Picture = product.Picture
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory

		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}

// Search product by category filter
func (s *productService) GetProductsByCategory(idCategory int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProductsByCategory(idCategory)
	var productsDto dto.ProductsDto
	
	if len(products) == 0 {
		return productsDto, e.NewBadRequestApiError("products not found")
	}
	
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id = product.Id
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Picture = product.Picture
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory

		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}

// Display random products by cantidad

func (s *productService) GetProductsByNumber(cantidad int) (dto.ProductsDto, e.ApiError) {
	var products model.Products = productCliente.GetProductsByNumber(cantidad)
	var productsDto dto.ProductsDto
	if len(products) == 0 {
		return productsDto, e.NewBadRequestApiError("products not found")
	}
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id = product.Id
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Picture = product.Picture
		productDto.Price = product.Price
		productDto.Stock = product.Stock
		productDto.IdCategory = product.IdCategory

		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}