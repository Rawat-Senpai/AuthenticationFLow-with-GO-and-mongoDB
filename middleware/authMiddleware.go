package middleware

import (
	"fmt"
	"github/rawat-senpai/helpers"
	"net/http"

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
