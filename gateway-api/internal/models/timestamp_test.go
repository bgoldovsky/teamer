package models_test

import (
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func TestToTime(t *testing.T) {
	stamp := &timestamp.Timestamp{Seconds: 10000000, Nanos: 10000}
	expected, _ := time.Parse("2006-01-02 15:04:05.000000", "1970-04-26 17:46:40.000010")

	act := models.ToTime(stamp)

	if act != expected {
		t.Errorf("expected %v, act %v", expected, act)
	}
}

func TestToTimestamp(t *testing.T) {
	expected := &timestamp.Timestamp{Seconds: 10000000, Nanos: 10000}
	actTime, _ := time.Parse("2006-01-02 15:04:05.000000", "1970-04-26 17:46:40.000010")

	act := models.ToTimestamp(actTime)

	if act.Seconds != expected.Seconds {
		t.Errorf("expected %v, act %v", expected.Seconds, act.Seconds)
	}

	if act.Nanos != expected.Nanos {
		t.Errorf("expected %v, act %v", expected.Nanos, act.Nanos)
	}
}
