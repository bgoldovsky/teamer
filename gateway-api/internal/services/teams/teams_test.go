package teams_test

import (
	"context"
	"testing"
	"time"

	"github.com/bgoldovsky/teamer/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/teamer/gateway-api/internal/models"
	teamsRepo "github.com/bgoldovsky/teamer/gateway-api/internal/repostiory/teams"
	. "github.com/bgoldovsky/teamer/gateway-api/internal/services/teams"
	"github.com/stretchr/testify/assert"
)

var team = &models.TeamView{
	ID:          777,
	Name:        "Dream Team",
	Description: "Best team ever",
	Slack:       "QWERTY",
	Crated:      time.Now().UTC(),
	Updated:     time.Now().UTC(),
}

func TestService_GetTeams(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	act, err := s.GetTeams(context.Background())

	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act[0].ID, team.ID)
	assert.Equal(t, act[0].Name, team.Name)
	assert.Equal(t, act[0].Description, team.Description)
	assert.Equal(t, act[0].Slack, team.Slack)
	assert.Equal(t, act[0].Crated, team.Crated)
	assert.Equal(t, act[0].Updated, team.Updated)
}

func TestService_RemoveTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	act, err := s.RemoveTeam(context.Background(), team.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, team.ID)
	assert.Equal(t, act.Msg, "successfully removed")
}

func TestService_UpdateTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	form := models.TeamForm{
		Name:        team.Name,
		Description: team.Description,
		Slack:       team.Slack,
	}

	act, err := s.UpdateTeam(context.Background(), team.ID, &form)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, team.ID)
	assert.Equal(t, act.Msg, "successfully updated")
}

func TestService_AddTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	form := models.TeamForm{
		Name:        team.Name,
		Description: team.Description,
		Slack:       team.Slack,
	}

	act, err := s.AddTeam(context.Background(), &form)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, team.ID)
	assert.Equal(t, act.Msg, "successfully created")
}

func newClientMock() *teams.Client {
	return teams.NewMock(
		team.ID,
		team.Name,
		team.Description,
		team.Slack,
		team.Crated,
		team.Updated)
}

func newRepoMock() *teamsRepo.RepositoryMock {
	return teamsRepo.NewMock(
		team.ID,
		team.Name,
		team.Description,
		team.Slack,
		team.Crated,
		team.Updated)
}
