package repository

import (
	"toko1/helper"
	"toko1/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	GetTransactions(userID uint) ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	err := r.db.Preload("Product").Preload("User").Create(&transaction).Error
	return transaction, helper.ReturnIfError(err)
}

func (r *transactionRepository) GetTransactions(userID uint) ([]entity.Transaction, error) {
	var (
		transactions []entity.Transaction
	)

	db := r.db
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}

	err := db.Find(&transactions).Preload("Product").Preload("User").Find(&transactions).Error

	return transactions, helper.ReturnIfError(err)
}
