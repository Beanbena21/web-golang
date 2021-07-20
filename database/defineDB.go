package database

import (
	"context"
	"errors"
	"time"

	"bean.com/web-service/models"
	"bean.com/web-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleRegister(register *models.Register) (*mongo.InsertOneResult, error) {
	user := models.User{
		ID:            primitive.NewObjectID(),
		User_name:     register.User_name,
		User_password: utils.GeneratePassword(register.User_password),
		Email:         register.Email,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, _ := HandleLogin(register.Email)
	if result.Email != "" {

		return nil, errors.New("email is exist")
	}

	value, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func HandleLogin(email string) (*models.User, error) {
	user := models.User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func HandleInfo(id string) (*models.User, error) {
	user := models.User{}
	err := collection.FindOne(context.TODO(), bson.M{"_id": utils.ObjToString(id)}).Decode(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}
