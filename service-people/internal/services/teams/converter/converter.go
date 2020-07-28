package converter

import (
	"time"

	v1 "github.com/bgoldovsky/teamer-bot/service-people/internal/generated/rpc/v1"
	"github.com/bgoldovsky/teamer-bot/service-people/internal/models"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ToModel(dto *v1.UpdateTeamRequest) *models.Team {
	if dto == nil {
		return nil
	}

	return &models.Team{
		ID:          dto.Id,
		Name:        dto.Name,
		Description: dto.Description,
		Slack:       dto.Slack,
	}
}

func ToDTO(model *models.Team) *v1.Team {
	if model == nil {
		return nil
	}

	return &v1.Team{
		Id:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Slack:       model.Slack,
		Created:     getTimestamp(model.Created),
		Updated:     getTimestamp(model.Updated),
	}
}

func getTimestamp(t time.Time) *timestamp.Timestamp {
	stamp, _ := ptypes.TimestampProto(t)
	return stamp
}
