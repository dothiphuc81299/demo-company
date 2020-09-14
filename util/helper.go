package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HelperParseStringToObjectID ...
func HelperParseStringToObjectID(val string) primitive.ObjectID {
	result, _ := primitive.ObjectIDFromHex(val)
	return result
}