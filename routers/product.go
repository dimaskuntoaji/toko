package routers

import (
	"toko1/handler"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, controller handler.ProductController) {
	router.GET("/products", controller.GetProducts)
}