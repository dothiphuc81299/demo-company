package util

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// HelperParseStringToObjectID ...
func HelperParseStringToObjectID(val string) (primitive.ObjectID, error){
	result, err := primitive.ObjectIDFromHex(val)
	return result,err
}

// HelperConvertTimeToTimestampProto ...
func HelperConvertTimestampProtoToTime(t *timestamppb.Timestamp) time.Time {
	result, _ := ptypes.Timestamp(t)
	return result
}
