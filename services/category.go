package services

import (
	"toko1/helper"
	"toko1/entity"
	"toko1/repository"
)

type CategoryServices interface {
	GetCategories() ([]entity.CategoryResponseGet, error)
}

type categoryServices struct {
	repository repository.CategoryRepository
}

func NewCategoryServices(repository repository.CategoryRepository) *categoryServices {
	return &categoryServices{repository}
}

//#MELIHAT DAFTAR KATEGORI BARANG

// GetCategories godoc
// @Summary      Product category
// @Description  View product list by product category
// @Accept       json
// @Produce      json
// @Success      200 {object} entity.CategoryResponseGet
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /categories [get]
func (s *categoryServices) GetCategories() ([]entity.CategoryResponseGet, error) {
	var (
		categories          []entity.Category
		categoriesResponses []entity.CategoryResponseGet
	)

	categories, err := s.repository.GetAllCategories()
	if err != nil {
		helper.PanicIfError(err)
	}

	for _, category := range categories {
		var categoryResponse entity.CategoryResponseGet

		categoryResponse.ID = category.ID
		categoryResponse.Type = category.Type
		categoryResponse.CreatedAt = category.CreatedAt
		categoryResponse.UpdatedAt = category.UpdatedAt
		var productResponses []entity.ProductResponse
		for _, product := range category.Product {
			productResponse := entity.ProductResponse{
				ID:         product.ID,
				Price:      product.Price,
				Title:      product.Title,
				Stock:      *product.Stock,
				CategoryID: product.CategoryID,
				CreatedAt:  product.CreatedAt,
			}
			productResponses = append(productResponses, productResponse)
		}
		categoryResponse.Product = productResponses

		categoriesResponses = append(categoriesResponses, categoryResponse)
	}

	return categoriesResponses, nil
}




