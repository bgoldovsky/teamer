package models_test

import (
	"testing"
	"time"

	v1 "github.com/bgoldovsky/dutyer/gateway-api/internal/generated/clients/teams/v1"
	. "github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/golang/protobuf/ptypes/wrappers"
)

var team = &v1.Team{
	Id:          777,
	Name:        "Dream Team",
	Description: "Best team ever",
	Slack:       "QWERTY",
	Created:     ToTimestamp(time.Now()),
	Updated:     ToTimestamp(time.Now()),
}

var person = &v1.Person{
	Id:         666,
	TeamId:     777,
	FirstName:  "Boris",
	MiddleName: &wrappers.StringValue{Value: "B"},
	LastName:   "Goldovsky",
	Birthday:   ToTimestamp(time.Now()),
	Email:      &wrappers.StringValue{Value: "bg@example.com"},
	Phone:      &wrappers.StringValue{Value: "+7903 333 33 33"},
	Slack:      "QWERTY",
	Role:       RoleBackEnd,
	IsActive:   true,
	Created:    ToTimestamp(time.Now()),
	Updated:    ToTimestamp(time.Now()),
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

	if act.Slack != team.Slack {
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

	if act[0].Slack != team.Slack {
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

func TestFromPersonReply(t *testing.T) {
	act := FromPersonReply(person)

	if act.ID != person.Id {
		t.Errorf("expected %v, act %v", act.ID, person.Id)
	}

	if act.TeamId != person.TeamId {
		t.Errorf("expected %v, act %v", act.TeamId, person.TeamId)
	}

	if act.FirstName != person.FirstName {
		t.Errorf("expected %v, act %v", act.FirstName, person.FirstName)
	}

	expMiddle := &person.MiddleName.Value
	if act.MiddleName != expMiddle {
		t.Errorf("expected %v, act %v", act.MiddleName, expMiddle)
	}

	if act.LastName != person.LastName {
		t.Errorf("expected %v, act %v", act.LastName, person.LastName)
	}

	expBirthday := ToTime(person.Birthday)
	if *act.Birthday != expBirthday {
		t.Errorf("expected %v, act %v", act.Birthday, expBirthday)
	}

	expPhone := &person.Phone.Value
	if act.Phone != expPhone {
		t.Errorf("expected %v, act %v", act.Phone, expPhone)
	}

	expEmail := &person.Email.Value
	if act.Email != expEmail {
		t.Errorf("expected %v, act %v", act.Email, expEmail)
	}

	if act.Slack != person.Slack {
		t.Errorf("expected %v, act %v", act.Slack, person.Slack)
	}

	expRole := int64(person.Role)
	if act.Role != expRole {
		t.Errorf("expected %v, act %v", act.Role, expRole)
	}

	if act.IsActive != person.IsActive {
		t.Errorf("expected %v, act %v", act.IsActive, person.IsActive)
	}

	expCreated := ToTime(person.Created)
	if act.Created != expCreated {
		t.Errorf("expected %v, act %v", act.ID, expCreated)
	}

	expUpdated := ToTime(person.Updated)
	if act.Updated != expUpdated {
		t.Errorf("expected %v, act %v", act.Updated, expUpdated)
	}
}

func TestFromPersonsReply(t *testing.T) {
	reply := &v1.GetPersonsReply{Persons: []*v1.Person{person}}
	act := FromPersonsReply(reply)

	if act[0].ID != person.Id {
		t.Errorf("expected %v, act %v", act[0].ID, person.Id)
	}

	if act[0].TeamId != person.TeamId {
		t.Errorf("expected %v, act %v", act[0].TeamId, person.TeamId)
	}

	if act[0].FirstName != person.FirstName {
		t.Errorf("expected %v, act %v", act[0].FirstName, person.FirstName)
	}

	expMiddle := &person.MiddleName.Value
	if act[0].MiddleName != expMiddle {
		t.Errorf("expected %v, act %v", act[0].MiddleName, expMiddle)
	}

	if act[0].LastName != person.LastName {
		t.Errorf("expected %v, act %v", act[0].LastName, person.LastName)
	}

	expBirthday := ToTime(person.Birthday)
	if *act[0].Birthday != expBirthday {
		t.Errorf("expected %v, act %v", act[0].Birthday, expBirthday)
	}

	expPhone := &person.Phone.Value
	if act[0].Phone != expPhone {
		t.Errorf("expected %v, act %v", act[0].Phone, expPhone)
	}

	expEmail := &person.Email.Value
	if act[0].Email != expEmail {
		t.Errorf("expected %v, act %v", act[0].Email, expEmail)
	}

	if act[0].Slack != person.Slack {
		t.Errorf("expected %v, act %v", act[0].Slack, person.Slack)
	}

	expRole := int64(person.Role)
	if act[0].Role != expRole {
		t.Errorf("expected %v, act %v", act[0].Role, expRole)
	}

	if act[0].IsActive != person.IsActive {
		t.Errorf("expected %v, act %v", act[0].IsActive, person.IsActive)
	}

	expCreated := ToTime(person.Created)
	if act[0].Created != expCreated {
		t.Errorf("expected %v, act %v", act[0].ID, expCreated)
	}

	expUpdated := ToTime(person.Updated)
	if act[0].Updated != expUpdated {
		t.Errorf("expected %v, act %v", act[0].Updated, expUpdated)
	}
}
