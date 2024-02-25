package handler

import (
	"net/http"
	"toko1/entity"
	"toko1/services"
	"toko1/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) UserController {
	return UserController{us}
}

func (handler *UserController) SignUp(c *gin.Context) {
	var registerInput entity.UserRegisterInput

	err := c.ShouldBindJSON(&registerInput)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed To register",
				utils.GetErrorData(err),
			),
		)
		return
	}

	user, err := handler.userService.Register(registerInput)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to register",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.NewResponse(
			http.StatusCreated,
			"Success to register",
			user,
		),
	)
}

func (handler *UserController) SignIn(c *gin.Context) {
	var loginInput entity.UserLoginInput

	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Something wrong with input",
				utils.GetErrorData(err),
			),
		)
		return
	}

	token, err := handler.userService.Login(loginInput)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to login",
				utils.GetErrorData(err),
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		utils.NewResponse(
			http.StatusOK,
			"Login Success",
			"Bearer "+token,
		),
	)
}