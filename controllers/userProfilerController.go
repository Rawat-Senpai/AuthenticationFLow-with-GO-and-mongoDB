package controllers

import (
	"context"
	"fmt"
	"github/rawat-senpai/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updateValues map[string]interface{}

		// Parse request body to extract values to update
		if err := c.BindJSON(&updateValues); err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Error: Invalid request body"))
			return
		}

		userId, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("User Id is not found in Header "))
			return
		}

		userIdString, ok := userId.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("User Id is not in string format"))
			return
		}

		// Check if updateValues is empty
		if len(updateValues) == 0 {
			c.JSON(http.StatusBadRequest, response.ErrorResponse("Error: No User values provided"))
			return
		}

		// Construct the update query
		update := bson.M{"$set": updateValues}
		userIdHex, err := primitive.ObjectIDFromHex(userIdString)
		if err != nil {
			// Handle error
			fmt.Println("Invalid note ID:", err)
			return
		}

		// Specify the filter to identify the note to update
		filter := bson.M{"_id": userIdHex}

		updateResult, err := userCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: Failed to update note details:"+err.Error()))
			return
		}

		// Check if any documents were matched and modified
		if updateResult.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, response.ErrorResponse("Error: No matching document found for update"))
			return
		}
		c.JSON(http.StatusOK, response.SuccessResponse("User Update successfully "))

	}
}
