package repository

import (
	"toko1/helper"
	"toko1/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(userID uint) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user entity.User) (entity.User, error) {
	err := ur.db.Create(&user).Error
	return user, helper.ReturnIfError(err)
}

func (ur *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := ur.db.Where("email = ?", email).Find(&user).Error
	return user, helper.ReturnIfError(err)
}

func (ur *userRepository) GetUserByID(userID uint) (entity.User, error) {
	var user entity.User
	err := ur.db.Where("id = ?", userID).First(&user).Error
	return user, helper.ReturnIfError(err)
}
