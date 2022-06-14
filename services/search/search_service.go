package services

import (
	searchCliente "ArquicturaSW/clients/search"
	dto "ArquicturaSW/dto"
)

type searchService struct{}

type searchServiceInterface interface {
	GetProductsBySearchParam(string) (dto.ProductsDto, error)
}

var (
	SearchService searchServiceInterface
)

func init() {
	SearchService = &searchService{}
}

func (s *searchService) GetProductsBySearchParam(ref string) (dto.ProductsDto, error) {

	products, err := searchCliente.GetProductsBySearchParam(ref)

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

	return productsDto, err
}