package services

import (
	categoryCliente "ArquicturaSW/clients/category"
	"ArquicturaSW/dto"
	"ArquicturaSW/model"
	e "ArquicturaSW/utils/errors"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategories() (dto.CategoriesDto, e.ApiError)
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {

	var categories model.Categories = categoryCliente.GetCategories()
	var categoriesDto dto.CategoriesDto

	if len(categories) == 0 {
		return categoriesDto, e.NewBadRequestApiError("Categories not found")
	}

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.IdCategory = category.Id
		categoryDto.Name = category.Name

		categoriesDto = append(categoriesDto, categoryDto)
	}

	return categoriesDto, nil
}
