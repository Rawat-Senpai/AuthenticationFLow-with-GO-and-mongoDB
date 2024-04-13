package main

import (
	"fmt"
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

	// notesRouter := gin.New()
	// notesRouter.Use(middleware.Authentication())

	// router.Use(middleware.Authentication())
	// routes.NotesRoutes(notesRouter)

	router.Run(":" + port)

}
