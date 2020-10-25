package converter

import (
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func TestToModel(t *testing.T) {
	middleName := wrappers.StringValue{Value: "B"}
	email := wrappers.StringValue{Value: "bg@example.com"}
	phone := wrappers.StringValue{Value: "+79033333333"}
	birthday := timestamp.Timestamp{Seconds: 10000}

	dto := &v1.UpdatePersonRequest{
		Id:         111,
		TeamId:     222,
		FirstName:  "Boris",
		MiddleName: &middleName,
		LastName:   "Goldovskiy",
		Birthday:   &birthday,
		Email:      &email,
		Phone:      &phone,
		Slack:      "QWERTY",
		Role:       2,
		IsActive:   true,
	}

	act := ToModel(dto)

	if act.ID != dto.Id {
		t.Errorf("expected: %v, act: %v", dto.Id, act.ID)
	}

	if act.TeamID != dto.TeamId {
		t.Errorf("expected: %v, act: %v", dto.TeamId, act.TeamID)
	}

	if act.FirstName != dto.FirstName {
		t.Errorf("expected: %v, act: %v", dto.FirstName, act.FirstName)
	}

	if *act.MiddleName != dto.MiddleName.Value {
		t.Errorf("expected: %v, act: %v", dto.MiddleName.Value, *act.MiddleName)
	}

	if act.LastName != dto.LastName {
		t.Errorf("expected: %v, act: %v", dto.LastName, act.LastName)
	}

	if actBirthday := act.Birthday.Unix(); actBirthday != dto.Birthday.Seconds {
		t.Errorf("expected: %v, act: %v", dto.Birthday.Seconds, actBirthday)
	}

	if *act.Email != dto.Email.Value {
		t.Errorf("expected: %v, act: %v", dto.Email.Value, *act.Email)
	}

	if *act.Phone != dto.Phone.Value {
		t.Errorf("expected: %v, act: %v", dto.Phone.Value, *act.Phone)
	}

	if act.Slack != dto.Slack {
		t.Errorf("expected: %v, act: %v", dto.Slack, act.Slack)
	}

	if act.Role != models.Role(dto.Role) {
		t.Errorf("expected: %v, act: %v", dto.Role, act.Role)
	}

	if act.IsActive != dto.IsActive {
		t.Errorf("expected: %v, act: %v", dto.IsActive, act.IsActive)
	}
}

func TestToDTO(t *testing.T) {
	middleName := "B"
	birthday := time.Unix(10000, 0)
	email := "bg@example.com"
	phone := "+79033333333"

	model := &models.Person{
		ID:         111,
		TeamID:     222,
		FirstName:  "Boris",
		MiddleName: &middleName,
		LastName:   "Goldovskiy",
		Birthday:   &birthday,
		Email:      &email,
		Phone:      &phone,
		Slack:      "QWERTY",
		Role:       2,
		IsActive:   true,
		Created:    time.Now(),
		Updated:    time.Now(),
	}

	act := ToDTO(model)

	if act.Id != model.ID {
		t.Errorf("expected: %v, act: %v", model.ID, act.Id)
	}

	if act.TeamId != model.TeamID {
		t.Errorf("expected: %v, act: %v", model.TeamID, act.TeamId)
	}

	if act.FirstName != model.FirstName {
		t.Errorf("expected: %v, act: %v", model.FirstName, act.FirstName)
	}

	if act.MiddleName.Value != *model.MiddleName {
		t.Errorf("expected: %v, act: %v", *model.MiddleName, act.MiddleName.Value)
	}

	if act.LastName != model.LastName {
		t.Errorf("expected: %v, act: %v", model.LastName, act.LastName)
	}

	if act.Email.Value != *model.Email {
		t.Errorf("expected: %v, act: %v", *model.Email, act.Email.Value)
	}

	if act.Phone.Value != *model.Phone {
		t.Errorf("expected: %v, act: %v", *model.Phone, act.Phone.Value)
	}

	if act.Slack != model.Slack {
		t.Errorf("expected: %v, act: %v", model.Slack, act.Slack)
	}

	if act.Role != v1.Role(model.Role) {
		t.Errorf("expected: %v, act: %v", model.Role, act.Role)
	}

	if act.IsActive != model.IsActive {
		t.Errorf("expected: %v, act: %v", model.IsActive, act.IsActive)
	}

	if act.Created.Seconds != model.Created.Unix() {
		t.Errorf("expected: %v, act: %v", model.Created.Unix(), act.Created.Seconds)
	}

	if act.Updated.Seconds != model.Updated.Unix() {
		t.Errorf("expected: %v, act: %v", model.Updated.Unix(), act.Updated.Seconds)
	}
}
