package routes

import (
	"github/rawat-senpai/controllers"

	"github.com/gin-gonic/gin"
)

// userRoutes Functions

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())

}
