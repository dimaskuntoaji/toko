package middlewares

import (
	"net/http"
	"toko1/helper"
	"toko1/utils"
	"strings"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	if authorizationHeader == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	tokenType, token := strings.Split(authorizationHeader, " ")[0], strings.Split(authorizationHeader, " ")[1]
	if tokenType == "" || tokenType != "Bearer" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	if token == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	userID, err := helper.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			utils.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				err.Error(),
			),
		)
		return
	}

	c.Set("userID", userID)

}