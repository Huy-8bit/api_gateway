package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimestampProtoToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

func TimeToTimestampProto(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func StringToObjectId(id string) primitive.ObjectID {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return objectId
}

func ObjectIdToString(objectId primitive.ObjectID) string {
	return objectId.Hex()
}
