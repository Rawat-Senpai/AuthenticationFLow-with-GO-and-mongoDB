package routes

import (
	"github/rawat-senpai/controllers"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(incomingRout *gin.Engine) {

	incomingRout.POST("/notes/add", controllers.CreateNoteHandler())
	incomingRout.GET("/notes", controllers.GetNotesHandler())

}
