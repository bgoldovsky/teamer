package teams

import (
	"context"
	"testing"
	"time"

	"github.com/bgoldovsky/teamer/service-people/internal/publisher"

	v1 "github.com/bgoldovsky/teamer/service-people/internal/generated/rpc/v1"
	"github.com/bgoldovsky/teamer/service-people/internal/models"
	"github.com/bgoldovsky/teamer/service-people/internal/repository/teams"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

var team = &models.Team{
	ID:          777,
	Name:        "Dream Team",
	Description: "Fail Better",
	Slack:       "QWERTY",
	Created:     time.Now(),
	Updated:     time.Now(),
}

func TestService_AddTeam(t *testing.T) {
	repo := teams.NewMock()
	repo.ConfigureSave(team.ID, team.Name, team.Description, team.Slack)
	pub := publisher.NewMock(eventTeamAdded, team.ID, topicTeams)
	service := New(repo, pub)

	request := &v1.AddTeamRequest{
		Name:        team.Name,
		Description: team.Description,
		Slack:       team.Slack,
	}

	act, err := service.AddTeam(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, team.ID, act.Id)
}

func TestService_UpdateTeam(t *testing.T) {
	repo := teams.NewMock()
	repo.ConfigureUpdate(team.ID, team.Name, team.Description, team.Slack, team.Created, team.Updated)
	// TODO: удалить placeholder
	pub := publisher.NewMock("", 0, "")
	service := New(repo, pub)

	request := &v1.UpdateTeamRequest{
		Id:          team.ID,
		Name:        team.Name,
		Description: team.Description,
		Slack:       team.Slack,
	}

	act, err := service.UpdateTeam(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, act)
}

func TestService_RemoveTeam(t *testing.T) {
	repo := teams.NewMock()
	repo.ConfigureRemove(team.ID)
	// TODO: удалить placeholder
	pub := publisher.NewMock("", 0, "")
	service := New(repo, pub)

	request := &v1.RemoveTeamRequest{Id: team.ID}

	act, err := service.RemoveTeam(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, act)
}

func TestService_GetTeams(t *testing.T) {
	repo := teams.NewMock()
	repo.ConfigureGet(team.ID, team.Name, team.Description, team.Slack, team.Created, team.Updated)
	// TODO: удалить placeholder
	pub := publisher.NewMock("", 0, "")
	service := New(repo, pub)

	created, _ := ptypes.TimestampProto(team.Created)
	updated, _ := ptypes.TimestampProto(team.Updated)

	request := &v1.GetTeamsRequest{
		Filter: &v1.TeamFilter{
			Ids: []int64{team.ID},
		},
	}

	act, err := service.GetTeams(context.Background(), request)

	assert.NoError(t, err)
	assert.NotEmpty(t, act.Teams)
	assert.Equal(t, team.ID, act.Teams[0].Id)
	assert.Equal(t, team.Name, act.Teams[0].Name)
	assert.Equal(t, team.Description, act.Teams[0].Description)
	assert.Equal(t, team.Slack, act.Teams[0].Slack)
	assert.Equal(t, created, act.Teams[0].Created)
	assert.Equal(t, updated, act.Teams[0].Updated)
}
