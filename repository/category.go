package repository

import (
	"toko1/helper"
	"toko1/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository interface {
	GetAllCategories() ([]entity.Category, error)
	GetDataByID(categoryID uint) (entity.Category, error)
	UpdateCategory(category entity.Category) (entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}


func (r *categoryRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Preload(clause.Associations).Find(&categories).Error

	return categories, helper.ReturnIfError(err)
}

func (r *categoryRepository) GetDataByID(categoryID uint) (entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ?", categoryID).First(&category).Error
	return category, helper.ReturnIfError(err)
}

func (r *categoryRepository) UpdateCategory(category entity.Category) (entity.Category, error) {
	err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&category).Error
	return category, helper.ReturnIfError(err)
}



