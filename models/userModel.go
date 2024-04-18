package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"UserId"`
	Name          *string            `bson:"name" json:"name" validate:"required,min=2,max=20"`
	Passowrd      *string            `bson:"password" json:"password" validate:"required"`
	Email         *string            `bson:"email" json:"email" validate:"required"`
	Token         *string            `bson:"token" json:"token"`
	Refresh_token *string            `bson:"refresh_token" json:"refresh_token"`
	User_id       string             `bson:"user_id" json:"user_id"`
	Created_at    time.Time          `bson:"created_at" json:"created_at"`
	Updated_at    time.Time          `bson:"updated_at" json:"updated_at"`
	Phone         *string            `bson:"phone" json:"phone"`
	Profile       string             `bson:"userProfile" json:"userProfile"`
	OTP           string             `bson:"otp" json:"otp"`
}
