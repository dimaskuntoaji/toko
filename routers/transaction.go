package routers

import (
	"toko1/handler"
	"toko1/middlewares"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(c *gin.Engine, controller handler.TransactionController) {
	c.POST("/checkout", middlewares.AuthMiddleware, controller.CreateTransaction)
	c.GET("/transactions", middlewares.AuthMiddleware, controller.GetUserTransactions)
}
