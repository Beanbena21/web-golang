package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User_name     string             `bson:"user_name" json:"user_name"`
	User_password string             `bson:"user_password" json:"user_password"`
	Email         string             `bson:"email" json:"email"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Login struct {
	Email         string `bson:"email" json:"email" binding:"required"`
	User_password string `bson:"user_password" json:"user_password" binding:"required"`
}

type Register struct {
	User_name     string `bson:"user_name" json:"user_name" binding:"required"`
	User_password string `bson:"user_password" json:"user_password" binding:"required"`
	Email         string `bson:"email" json:"email" binding:"required"`
}

type Claims struct {
	Email string `bson:"email" json:"email"`
	jwt.StandardClaims
}

type Id struct {
	ID string `uri:"id" binding:"required"`
}
