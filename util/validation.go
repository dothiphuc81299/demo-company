package util

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidationObjectID ...
func ValidationObjectID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}
	return objectID, err
}
