package main

import (
	"fmt"
	"github/rawat-senpai/middleware"
	"github/rawat-senpai/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("set up is started ")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	//API  2

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api -1 "})

	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)

}
