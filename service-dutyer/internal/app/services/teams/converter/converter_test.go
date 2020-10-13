package converter

import (
	"testing"
	"time"

	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
)

func TestToDTO(t *testing.T) {
	model := &models.Team{
		ID:          123,
		Name:        "Dream Team",
		Description: "Best for everything",
		Slack:       "QWERTY",
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	act := ToDTO(model)

	if act.Id != model.ID {
		t.Errorf("expected %v, act %v", model.ID, act.Id)
	}

	if act.Name != model.Name {
		t.Errorf("expected %v, act %v", model.Name, act.Name)
	}

	if act.Description != model.Description {
		t.Errorf("expected %v, act %v", model.Description, act.Description)
	}

	if act.Slack != model.Slack {
		t.Errorf("expected %v, act %v", model.Slack, act.Slack)
	}

	if act.Created.Seconds != model.Created.Unix() {
		t.Errorf("expected %v, act %v", model.Created.Unix(), act.Created.Seconds)
	}
}

func TestToModel(t *testing.T) {
	dto := &v1.UpdateTeamRequest{
		Id:          123,
		Name:        "Dream Team",
		Description: "Best for everything",
		Slack:       "QWERTY",
	}

	act := ToModel(dto)

	if act.ID != dto.Id {
		t.Errorf("expected %v, act %v", dto.Id, act.ID)
	}

	if act.Name != dto.Name {
		t.Errorf("expected %v, act %v", dto.Name, act.Name)
	}

	if act.Description != dto.Description {
		t.Errorf("expected %v, act %v", dto.Description, act.Description)
	}

	if act.Slack != dto.Slack {
		t.Errorf("expected %v, act %v", dto.Slack, act.Slack)
	}
}
