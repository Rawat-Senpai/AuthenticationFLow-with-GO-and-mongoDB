package controllers

import (
	"context"
	"fmt"
	"github/rawat-senpai/models"
	"github/rawat-senpai/response"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func generateRandomString() string {
	// Seed the random number generator with the current timestamp
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 100000 and 999999 (inclusive)
	randomNumber := rand.Intn(900000) + 100000

	// Convert the random number to a string
	randomString := strconv.Itoa(randomNumber)

	return randomString
}

func sendOTP(email, otp string) error {
	auth := smtp.PlainAuth(
		"",
		"shobhitrawat84@gmail.com",
		"pnlgneiiyjfccpbt",
		"smtp.gmail.com",
	)

	msg := "Subject: OTP for Password Reset\r\n\r\nYour OTP for password reset is: " + otp

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"shobhitrawat84@gmail.com",
		[]string{email},
		[]byte(msg),
	)

	return err
}

func ForgotPasswordSendOtp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		var userModel models.AuthenticationModel
		if err := c.BindJSON(&userModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the user with the provided email exists
		var foundUser models.AuthenticationModel
		err := userCollection.FindOne(ctx, bson.M{"email": userModel.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: This email does not exist"))
			return
		}

		fmt.Println("Found User", foundUser)
		// Generate a random OTP
		randomString := generateRandomString()

		// Send email with the OTP
		err = sendOTP(userModel.Email, randomString)
		if err != nil {

			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: Failed to send OTP"+err.Error()))
			return
		}

		// Update the OTP value in the database
		update := bson.M{"$set": bson.M{"otp": randomString}}
		filter := bson.M{"email": userModel.Email}
		_, updateErr := userCollection.UpdateOne(ctx, filter, update)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: Failed to update OTP"))
			return
		}

		c.JSON(http.StatusOK, response.SuccessResponse("OTP sent successfully"))
	}
}

func ConfirmOtp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		var userModel models.AuthenticationModel
		if err := c.BindJSON(&userModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the user with the provided email exists
		var foundUser models.AuthenticationModel
		err := userCollection.FindOne(ctx, bson.M{"email": userModel.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: This email does not exist"))
			return
		}

		// Update the OTP value in the database
		update := bson.M{"$set": bson.M{"otp": ""}}
		filter := bson.M{"email": userModel.Email}
		_, updateErr := userCollection.UpdateOne(ctx, filter, update)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: Failed to update OTP"))
			return
		}

		c.JSON(http.StatusOK, response.SuccessResponse("OTP sent successfully"))

	}
}
