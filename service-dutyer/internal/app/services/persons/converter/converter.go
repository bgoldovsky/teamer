package converter

import (
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func ToModel(dto *v1.UpdatePersonRequest) *models.Person {
	if dto == nil {
		return nil
	}

	person := &models.Person{
		ID:        dto.Id,
		TeamID:    dto.TeamId,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Birthday:  toTimestamp(dto.Birthday),

		Slack: dto.Slack,
		Role:  models.Role(dto.Role),

		IsActive: dto.IsActive,
	}

	if dto.MiddleName != nil {
		person.MiddleName = &dto.MiddleName.Value
	}

	if dto.Email != nil {
		person.Email = &dto.Email.Value
	}

	if dto.Phone != nil {
		person.Phone = &dto.Phone.Value
	}

	return person
}

func ToDTO(model *models.Person) *v1.Person {
	if model == nil {
		return nil
	}

	person := &v1.Person{
		Id:        model.ID,
		TeamId:    model.TeamID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Slack:     model.Slack,
		Role:      v1.Role(model.Role),
		DutyOrder: model.DutyOrder,
		IsActive:  model.IsActive,
		Created:   getTimestamp(&model.Created),
		Updated:   getTimestamp(&model.Updated),
	}

	if model.MiddleName != nil {
		person.MiddleName = &wrappers.StringValue{Value: *model.MiddleName}
	}

	if model.Birthday != nil {
		person.Birthday = getTimestamp(model.Birthday)
	}

	if model.Email != nil {
		person.Email = &wrappers.StringValue{Value: *model.Email}
	}

	if model.Phone != nil {
		person.Phone = &wrappers.StringValue{Value: *model.Phone}
	}

	return person
}

func getTimestamp(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}

	stamp, _ := ptypes.TimestampProto(*t)
	return stamp
}

func toTimestamp(t *timestamp.Timestamp) *time.Time {
	if t == nil {
		return nil
	}

	stamp, _ := ptypes.Timestamp(t)
	return &stamp
}
