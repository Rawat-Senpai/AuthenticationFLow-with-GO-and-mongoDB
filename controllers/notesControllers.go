package controllers

import (
	"context"
	"fmt"
	"github/rawat-senpai/database"
	"github/rawat-senpai/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var noteCollection *mongo.Collection = database.OpenCollection(database.Client, "notes")

func CreateNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var note models.Notes
		token := c.Request.Header.Get("token")
		fmt.Printf(token)
		if err := c.BindJSON(&note); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Authorization Header Not Found"})
			return
		}

		userId, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User id not found in context"})
			return
		}

		userIdString, ok := userId.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User id is not a string"})
			return
		}

		// Associate the note with the user
		note.CreatedBy = userIdString
		// Associate the note with the user
		// note.CreatedBy = userId.(string)

		result, err := noteCollection.InsertOne(context.Background(), note)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}

func GetNotesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User id not found in context"})
			return
		}

		cursor, err := noteCollection.Find(context.Background(), bson.M{"createdBy": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive the notes "})
			fmt.Println("Error querying database:", err)
			return
		}

		defer cursor.Close(context.Background())

		var notes []models.Notes
		// Create a new slice to store filtered notes

		if err := cursor.All(context.Background(), &notes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
			fmt.Println("Error decoding notes:", err)
			return
		}

		fmt.Println("user tokenoutside for loop ", userId)
		for _, note := range notes {
			fmt.Println("user token", note.CreatedBy)

		}
		// Return the retrieved notes
		c.JSON(http.StatusOK, notes)

	}
}

func GetAllNotesNotesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		cursor, err := noteCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive the notes "})
			fmt.Println("Error querying database:", err)
			return
		}

		defer cursor.Close(context.Background())

		var notes []models.Notes

		if err := cursor.All(context.Background(), &notes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
			fmt.Println("Error decoding notes:", err)
			return
		}

		userId, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User id not found in context"})
			return
		}

		for _, note := range notes {

			fmt.Println("user id:", note.CreatedBy)

			if userId == note.CreatedBy {

				fmt.Println("is true", "true")
			}

		}

		c.JSON(http.StatusOK, notes)

	}
}
