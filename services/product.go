package services

import (
	"toko1/helper"
	"toko1/entity"
	"toko1/repository"

)

type ProductServices interface {
	GetProducts() ([]entity.ProductResponse, error)
}

type productServices struct {
	repository         repository.ProductRepository
	categoryRepository repository.CategoryRepository
}

func NewProductServices(repository repository.ProductRepository, categoryRepository repository.CategoryRepository) *productServices {
	return &productServices{repository, categoryRepository}
}

//#MELIHAT DAFTAR PRODUCT

// Product details godoc
// @Summary      Product details
// @Description  Product details
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /products [get]
func (s *productServices) GetProducts() ([]entity.ProductResponse, error) {
	var (
		products         []entity.Product
		productResponses []entity.ProductResponse
	)

	products, err := s.repository.GetProducts()

	for _, product := range products {
		var productResponse entity.ProductResponse

		productResponse.ID = product.ID
		productResponse.Title = product.Title
		productResponse.Price = product.Price
		productResponse.Stock = *product.Stock
		productResponse.CategoryID = product.CategoryID
		productResponse.CreatedAt = product.CreatedAt

		productResponses = append(productResponses, productResponse)
	}

	return productResponses, helper.ReturnIfError(err)
}

