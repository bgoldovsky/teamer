package models_test

import (
	"testing"
	"time"

	v1 "github.com/bgoldovsky/teamer-bot/gateway-api/internal/generated/clients/people/v1"
	. "github.com/bgoldovsky/teamer-bot/gateway-api/internal/models"
)

var team = &v1.Team{
	Id:          777,
	Name:        "Dream Team",
	Description: "Best team ever",
	Slack:       "QWERTY",
	Created:     ToTimestamp(time.Now()),
	Updated:     ToTimestamp(time.Now()),
}

func TestNewStatusView(t *testing.T) {
	var id int64 = 777
	msg := "success"

	act := NewStatusView(id, msg)

	if act.ID != id {
		t.Errorf("expected %v, act %v", id, act.ID)
	}

	if act.Msg != msg {
		t.Errorf("expected %v, act %v", msg, act.Msg)
	}
}

func TestFromTeamReply(t *testing.T) {
	act := FromTeamReply(team)

	if act.ID != team.Id {
		t.Errorf("expected %v, act %v", act.ID, team.Id)
	}

	if act.Name != team.Name {
		t.Errorf("expected %v, act %v", act.Name, team.Name)
	}

	if act.ID != team.Id {
		t.Errorf("expected %v, act %v", act.Description, team.Description)
	}

	if act.ID != team.Id {
		t.Errorf("expected %v, act %v", act.Slack, team.Slack)
	}

	expCreated := ToTime(team.Created)
	if act.Crated != expCreated {
		t.Errorf("expected %v, act %v", act.ID, expCreated)
	}

	expUpdated := ToTime(team.Updated)
	if act.Updated != expUpdated {
		t.Errorf("expected %v, act %v", act.Updated, expUpdated)
	}
}

func TestFromTeamsReply(t *testing.T) {
	reply := &v1.GetTeamsReply{Teams: []*v1.Team{team}}
	act := FromTeamsReply(reply)

	if act[0].ID != team.Id {
		t.Errorf("expected %v, act %v", act[0].ID, team.Id)
	}

	if act[0].Name != team.Name {
		t.Errorf("expected %v, act %v", act[0].Name, team.Name)
	}

	if act[0].ID != team.Id {
		t.Errorf("expected %v, act %v", act[0].Description, team.Description)
	}

	if act[0].ID != team.Id {
		t.Errorf("expected %v, act %v", act[0].Slack, team.Slack)
	}

	expCreated := ToTime(team.Created)
	if act[0].Crated != expCreated {
		t.Errorf("expected %v, act %v", act[0].ID, expCreated)
	}

	expUpdated := ToTime(team.Updated)
	if act[0].Updated != expUpdated {
		t.Errorf("expected %v, act %v", act[0].Updated, expUpdated)
	}
}
