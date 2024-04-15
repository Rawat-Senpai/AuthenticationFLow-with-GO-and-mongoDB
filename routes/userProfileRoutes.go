package routes

import (
	"github/rawat-senpai/controllers"
	"github/rawat-senpai/middleware"

	"github.com/gin-gonic/gin"
)

func UserProfileRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.POST("/usersProfile/updateUser", controllers.UpdateUserProfile())

}
