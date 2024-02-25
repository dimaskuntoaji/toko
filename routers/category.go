package routers

import (
	"toko1/handler"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine, controller handler.CategoryController) {
	router.GET("/categories", controller.GetCategories)
}