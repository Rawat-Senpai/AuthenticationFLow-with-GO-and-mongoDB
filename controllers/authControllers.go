package controllers

import (
	"context"
	"fmt"
	"github/rawat-senpai/models"
	"github/rawat-senpai/response"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ForgotPasswordSendOtp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		var userModel models.AuthenticationModel
		if err := c.BindJSON(&userModel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		emailErr := userCollection.FindOne(ctx, bson.M{"email": userModel.Email})

		if emailErr != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error : This email does not exists "))
			return
		}

		randomString := generateRandomString()

		auth := smtp.PlainAuth(
			"",
			"shobhitrawat84@gmail.com",
			os.Getenv("GMAIL_PASSWORD"),
			"smtp.gmail.com",
		)

		msg := "Subject:- The verification otp is " + randomString

		fmt.Println(userModel.Email)

		err := smtp.SendMail(
			"smtp.gmail.com:587",
			auth,
			"shobhitrawat84@gmail.com",
			[]string{userModel.Email},
			[]byte(msg),
		)

		if err != nil {
			fmt.Println(err)
		}
		update := bson.M{"$set": bson.M{"otp": randomString}}
		filter := bson.M{"email": userModel.Email}

		_, UpdateErr := userCollection.UpdateOne(context.Background(), filter, update)

		if UpdateErr != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error :  "+err.Error()))
			return
		}

		// newErr:=	 userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse("Error: "+err.Error()))
			return
		}

		c.JSON(http.StatusOK, response.SuccessResponse("Password send successfully "))

	}
}

func generateRandomString() string {
	// Seed the random number generator with the current timestamp
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 100000 and 999999 (inclusive)
	randomNumber := rand.Intn(900000) + 100000

	// Convert the random number to a string
	randomString := strconv.Itoa(randomNumber)

	return randomString
}
