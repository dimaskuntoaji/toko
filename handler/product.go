package handler

import (
	"net/http"
	"toko1/services"
	"toko1/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductServices
}

func NewProductController(productService services.ProductServices) ProductController {
	return ProductController{productService}
}

func (handler *ProductController) GetProducts(c *gin.Context) {
	products, err := handler.productService.GetProducts()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get data",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Successfully get data",
			products,
		),
	)
}



