package utils 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidationObjectID ...
func ValidationObjectID(id string) (primitive.ObjectID,error) {
	result, err := primitive.ObjectIDFromHex(id)
	return result,err
}
