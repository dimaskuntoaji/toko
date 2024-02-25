package routers

import (
	"toko1/handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, controller handler.UserController) {
	router.POST("users/register", controller.SignUp)
	router.POST("users/login", controller.SignIn)
}