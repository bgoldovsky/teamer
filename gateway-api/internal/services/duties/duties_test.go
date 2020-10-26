package duties

import (
	"context"
	"testing"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/duties"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/stretchr/testify/assert"
)

var (
	duty = &models.DutyView{
		TeamId:    888,
		PersonId:  777,
		FirstName: "Boris",
		LastName:  "B",
		Slack:     "QWERTY",
		Channel:   "YTREWQ",
		Order:     5,
	}

	secondPersonID int64 = 999
	count          int64 = 50
)

func TestService_Swap(t *testing.T) {
	client := newClientMock()
	s := New(client)

	act, err := s.Swap(context.Background(), duty.TeamId, duty.PersonId, secondPersonID)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, duty.TeamId)
	assert.Equal(t, act.Msg, "successfully swapped")
}

func TestService_Assign(t *testing.T) {
	client := newClientMock()
	s := New(client)

	act, err := s.Assign(context.Background(), duty.TeamId, duty.PersonId)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.ID, duty.TeamId)
	assert.Equal(t, act.Msg, "successfully assigned")
}

func TestService_GetCurrentDuty(t *testing.T) {
	client := newClientMock()
	s := New(client)

	act, err := s.GetCurrentDuty(context.Background(), duty.TeamId)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act.PersonId, duty.PersonId)
	assert.Equal(t, act.TeamId, duty.TeamId)
	assert.Equal(t, act.FirstName, duty.FirstName)
	assert.Equal(t, act.LastName, duty.LastName)
	assert.Equal(t, act.Order, duty.Order)
	assert.Equal(t, act.Slack, duty.Slack)
}

func TestService_GetDuties(t *testing.T) {
	client := newClientMock()
	s := New(client)

	act, err := s.GetDuties(context.Background(), duty.TeamId, count)
	assert.NoError(t, err)
	assert.NotEmpty(t, act)
	assert.Equal(t, act[0].PersonId, duty.PersonId)
	assert.Equal(t, act[0].TeamId, duty.TeamId)
	assert.Equal(t, act[0].FirstName, duty.FirstName)
	assert.Equal(t, act[0].LastName, duty.LastName)
	assert.Equal(t, act[0].Order, duty.Order)
	assert.Equal(t, act[0].Slack, duty.Slack)
}

func newClientMock() *duties.Client {
	return duties.NewMock(
		duty.TeamId,
		duty.PersonId,
		secondPersonID,
		count,
		duty.FirstName,
		duty.LastName,
		duty.Slack,
		duty.Channel,
		duty.Order)
}
