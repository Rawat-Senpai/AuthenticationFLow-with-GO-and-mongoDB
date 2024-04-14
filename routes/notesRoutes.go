package routes

import (
	"github/rawat-senpai/controllers"
	"github/rawat-senpai/middleware"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(incomingRout *gin.Engine) {
	incomingRout.Use(middleware.Authentication())
	incomingRout.POST("/notes/add", controllers.CreateNoteHandler())
	incomingRout.GET("/notes", controllers.GetNotesHandler())

}
