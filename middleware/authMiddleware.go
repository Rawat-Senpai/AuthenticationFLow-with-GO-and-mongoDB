package middleware

import (
	"fmt"
	"github/rawat-senpai/helpers"
	"github/rawat-senpai/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorize users

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		reqToken := c.Request.Header.Get("Authorization")

		if reqToken == "" {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: "+"Authorization Token Required"))
			return
		}

		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) < 2 {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("error"+"Authorization Token is invalid"))
		}

		reqToken = splitToken[1]

		fmt.Printf("final token passwd in bearer" + reqToken)

		claims, err := helpers.ValidateToken(reqToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: "+err))
			c.Abort()
			return
		}
		c.Set("uid", claims.Uid)

		c.Next()

	}

}
