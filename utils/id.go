package utils

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjToString(id string) primitive.ObjectID {
	result, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	return result
}
