package persons

import (
	"context"
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/persons"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	personsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/persons"
	"github.com/stretchr/testify/assert"
)

var person = &models.PersonView{
	ID:        777,
	TeamId:    888,
	FirstName: "Boris",
	LastName:  "B",
	Slack:     "QWERTY",
	Role:      2,
	Created:   time.Now().UTC(),
	Updated:   time.Now().UTC(),
}

func TestService_GetTeams(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	act, err := s.GetPersons(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act[0].ID, person.ID)
	assert.Equal(t, act[0].TeamId, person.TeamId)
	assert.Equal(t, act[0].FirstName, person.FirstName)
	assert.Equal(t, act[0].LastName, person.LastName)
	assert.Equal(t, act[0].Slack, person.Slack)
	assert.Equal(t, act[0].Role, person.Role)
	assert.Equal(t, act[0].Created, person.Created)
	assert.Equal(t, act[0].Updated, person.Updated)
}

func TestService_RemoveTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	act, err := s.RemovePerson(context.Background(), person.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, person.ID)
	assert.Equal(t, act.Msg, "successfully removed")
}

func TestService_UpdateTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	form := models.PersonForm{
		TeamId:    person.TeamId,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Slack:     person.Slack,
		Role:      person.Role,
		IsActive:  person.IsActive,
	}

	act, err := s.UpdatePerson(context.Background(), person.ID, &form)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, person.ID)
	assert.Equal(t, act.Msg, "successfully updated")
}

func TestService_AddTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := New(client, repo)

	form := models.PersonForm{
		TeamId:    person.TeamId,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Slack:     person.Slack,
		Role:      person.Role,
		IsActive:  person.IsActive,
	}

	act, err := s.AddPerson(context.Background(), &form)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, person.ID)
	assert.Equal(t, act.Msg, "successfully created")
}

func newClientMock() *persons.Client {
	return persons.NewMock(
		person.ID,
		&models.PersonForm{
			TeamId:    person.TeamId,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Slack:     person.Slack,
			Role:      person.Role,
			IsActive:  person.IsActive,
		},
		person.Created,
		person.Updated)
}

func newRepoMock() *personsRepo.RepositoryMock {
	return personsRepo.NewMock(
		person.ID,
		person.TeamId,
		person.FirstName,
		person.LastName,
		person.Slack,
		person.Role,
		person.Created,
		person.Updated)
}
