package services

import (
	"errors"
	"toko1/helper"
	"toko1/entity"
	"toko1/repository"
)

type TransactionService interface {
	CreateTransaction(input entity.TransactionInput, userID uint) (entity.TransactionPostResponse, error)
	GetTransactions(userID uint) ([]entity.UserTransactionResponse, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	productRepository     repository.ProductRepository
	userRepository        repository.UserRepository
	categoryRepository    repository.CategoryRepository
}

func NewTransactionService(
	transactionRepository repository.TransactionRepository,
	productRepository repository.ProductRepository,
	userRepository repository.UserRepository,
	categoryRepository repository.CategoryRepository,
) *transactionService {
	return &transactionService{
		transactionRepository,
		productRepository,
		userRepository,
		categoryRepository,
	}
}

//#CHECKOUT BARANG

// CreateTransaction godoc
// @Summary      Checkout
// @Description  Checkout product list in shopping cart
// @Tags         User
// @Accept       json
// @Param        request body entity.TransactionInput true "Payload Body [RAW]"
// @Produce      json
// @Success      200 {object} entity.Transaction
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /checkout [post]
// @Security BearerAuth
func (s *transactionService) CreateTransaction(input entity.TransactionInput, userID uint) (entity.TransactionPostResponse, error) {
	var (
		transactionResponse entity.TransactionPostResponse
		transaction         entity.Transaction
	)

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return transactionResponse, err
	}
	product, err := s.productRepository.GetDataByID(uint(input.ProductID))
	if err != nil {
		return transactionResponse, err
	}
	category, err := s.categoryRepository.GetDataByID(product.CategoryID)
	if err != nil {
		return transactionResponse, err
	}

	if *product.Stock < input.Quantity {
		return transactionResponse, errors.New("product is not available")
	}

	totalPrice := product.Price * input.Quantity
	category, err = s.categoryRepository.UpdateCategory(category)
	if err != nil {
		return transactionResponse, err
	}

	stock := *product.Stock - input.Quantity
	*product.Stock = stock
	product, err = s.productRepository.UpdateProduct(product)
	if err != nil {
		return transactionResponse, err
	}

	transaction.UserID = user.ID
	transaction.ProductID = product.ID
	transaction.Quantity = input.Quantity
	transaction.TotalPrice = totalPrice

	transaction, err = s.transactionRepository.CreateTransaction(transaction)
	transactionResponse = entity.TransactionPostResponse{
		TotalPrice:   transaction.TotalPrice,
		Quantity:     transaction.Quantity,
		ProductTitle: product.Title,
	}

	return transactionResponse, helper.ReturnIfError(err)
}

//# MELIHAT DAFTAR CHECKOUT

// GetTransactions godoc
// @Summary      Payment transactions
// @Description  List of payment transactions
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200 {object} entity.UserTransactionResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /transactions [get]
// @Security BearerAuth
func (s *transactionService) GetTransactions(userID uint) ([]entity.UserTransactionResponse, error) {
	transactions, err := s.transactionRepository.GetTransactions(userID)
	var transactionResponses []entity.UserTransactionResponse
	for _, transaction := range transactions {
		transactionResponse := entity.UserTransactionResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: entity.ProductResponse{
				ID:         transaction.Product.ID,
				Title:      transaction.Product.Title,
				Price:      transaction.Product.Price,
				Stock:      *transaction.Product.Stock,
				CategoryID: transaction.Product.CategoryID,
				CreatedAt:  transaction.Product.CreatedAt,
			},
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, helper.ReturnIfError(err)
}
