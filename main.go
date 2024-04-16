package main

import (
	"fmt"
	"github/rawat-senpai/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("set up is started ")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env files")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.NotesRoutes(router)
	routes.UserProfileRoutes(router)

	router.Run(":" + port)

}

// pnlg neii yjfc cpbt
