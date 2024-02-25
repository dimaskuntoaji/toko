package repository

import (
	"toko1/helper"
	"toko1/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]entity.Product, error)
	GetDataByID(productID uint) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}


func (r *productRepository) GetProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, helper.ReturnIfError(err)
}

func (pr *productRepository) GetDataByID(productID uint) (entity.Product, error) {
	var product entity.Product
	err := pr.db.Preload("Category").Where("id = ?", productID).Find(&product).Error
	return product, helper.ReturnIfError(err)
}

func (pr *productRepository) UpdateProduct(product entity.Product) (entity.Product, error) {
	err := pr.db.Where("id = ?", product.ID).Updates(&product).Error
	return product, helper.ReturnIfError(err)
}