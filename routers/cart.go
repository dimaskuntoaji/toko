package routers

import (
	"toko1/handler"
	"toko1/middlewares"

	"github.com/gin-gonic/gin"
)

func CartRoute(c *gin.Engine, controller handler.CartController) {
	c.POST("/cart", middlewares.AuthMiddleware, controller.CreateCart)
	c.GET("/cart", middlewares.AuthMiddleware, controller.GetUserCarts)
	c.DELETE("/cart/:id", middlewares.AuthMiddleware, controller.DeleteCart)
}
