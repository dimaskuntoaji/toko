package handler

import (
	"net/http"
	"toko1/entity"
	"toko1/services"
	"toko1/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	s services.CartService
}

func NewCartController(s services.CartService) CartController {
	return CartController{s}
}


//#RESPON JSON MENAMBAHKAN BARANG KE KERANJANG
func (handler *CartController) CreateCart(c *gin.Context) {
	var input entity.CartInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create cart",
				utils.GetErrorData(err),
			),
		)
		return
	}

	userID := c.GetUint("userID")

	cart, err := handler.s.CreateCart(input, userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to create cart",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.NewResponse(
			http.StatusCreated,
			"Success to create cart",
			cart,
		),
	)
}

//#RESPON JSON MELIHAT USER BARANG KE KERANJANG
func (handler *CartController) GetUserCarts(c *gin.Context) {
	userID := c.GetUint("userID")

	carts, err := handler.s.GetCarts(userID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get carts",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Success to get carts",
			carts,
		),
	)
}

//#RESPON JSON MENGHAPUS BARANG KE KERANJANG
func (handler *CartController) DeleteCart(c *gin.Context) {
	var cartIDRaw = c.Param("id")

	cartID, err := strconv.Atoi(cartIDRaw)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Parameter must be a valid ID",
				utils.GetErrorData(err),
			),
		)
		return
	}

	err = handler.s.DeleteCart(uint(cartID))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to delete data",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Successfully delete data",
			nil,
		),
	)
}
