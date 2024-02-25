package handler

import (
	"net/http"
	"toko1/entity"
	"toko1/services"
	"toko1/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	s services.TransactionService
}

func NewTransactionController(s services.TransactionService) TransactionController {
	return TransactionController{s}
}

func (handler *TransactionController) CreateTransaction(c *gin.Context) {
	var input entity.TransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create transaction",
				utils.GetErrorData(err),
			),
		)
		return
	}

	userID := c.GetUint("userID")

	transaction, err := handler.s.CreateTransaction(input, userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create transaction",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.NewResponse(
			http.StatusCreated,
			"Success to create transaction",
			transaction,
		),
	)
}

func (handler *TransactionController) GetUserTransactions(c *gin.Context) {
	userID := c.GetUint("userID")

	transactions, err := handler.s.GetTransactions(userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get transactions",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Success to get transactions",
			transactions,
		),
	)
}

