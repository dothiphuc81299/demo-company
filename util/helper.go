package util

import (
	"github.com/golang/protobuf/ptypes"
	"time"
	"google.golang.org/protobuf/types/known/timestamppb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HelperParseStringToObjectID ...
func HelperParseStringToObjectID(val string) primitive.ObjectID {
	result, _ := primitive.ObjectIDFromHex(val)
	return result
}

// HelperConvertTimeToTimestampProto ...
func HelperConvertTimestampProtoToTime(t *timestamppb.Timestamp) (time.Time) {
	result, _ := ptypes.Timestamp(t)
	return result
}
