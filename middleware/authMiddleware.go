package middleware

import (
	"fmt"
	"github/rawat-senpai/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorize users

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authentication header provided")})
			c.Abort()
			return
		}

		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		fmt.Printf("final token passwd in bearer" + reqToken)

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("uid", claims.Uid)

		c.Next()

	}

}
