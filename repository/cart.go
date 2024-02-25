package repository

import (
	"toko1/helper"
	"toko1/entity"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart entity.Cart) (entity.Cart, error)
	GetCarts(userID uint) ([]entity.Cart, error)
	GetDataByID(cartID uint) (entity.Cart, error)
	DeleteCart(cart entity.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(cart entity.Cart) (entity.Cart, error) {
	err := r.db.Preload("Product").Preload("User").Create(&cart).Error
	return cart, helper.ReturnIfError(err)
}

func (r *cartRepository) GetCarts(userID uint) ([]entity.Cart, error) {
	var (
		carts []entity.Cart
	)

	db := r.db
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}

	err := db.Find(&carts).Preload("Product").Preload("User").Find(&carts).Error

	return carts, helper.ReturnIfError(err)
}

func (r *cartRepository) DeleteCart(cart entity.Cart) error {
	err := r.db.Delete(&cart).Error
	return helper.ReturnIfError(err)
}

func (r *cartRepository) GetDataByID(cartID uint) (entity.Cart, error) {
	var cart entity.Cart
	err := r.db.Where("id = ?", cartID).First(&cart).Error
	return cart, helper.ReturnIfError(err)
}
