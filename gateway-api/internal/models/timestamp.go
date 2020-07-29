package models

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ToTime(stamp *timestamp.Timestamp) time.Time {
	t, _ := ptypes.Timestamp(stamp)
	return t
}

func ToTimestamp(t time.Time) *timestamp.Timestamp {
	stamp, _ := ptypes.TimestampProto(t)
	return stamp
}
