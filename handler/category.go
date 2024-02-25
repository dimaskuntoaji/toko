package handler

import (
	"net/http"
	"toko1/services"
	"toko1/utils"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.CategoryServices
}

func NewCategoryController(service services.CategoryServices) CategoryController {
	return CategoryController{service}
}


func (handler *CategoryController) GetCategories(c *gin.Context) {
	categories, err := handler.categoryService.GetCategories()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get all data",
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
			categories,
		),
	)
}



