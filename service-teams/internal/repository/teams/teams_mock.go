package teams

import (
	"context"
	"time"

	v1 "github.com/bgoldovsky/dutyer/service-teams/internal/generated/rpc/v1"
	"github.com/bgoldovsky/dutyer/service-teams/internal/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func NewMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (m *RepositoryMock) Save(_ context.Context, team *models.Team) (*models.Team, error) {
	args := m.Called(team)
	res, _ := args.Get(0).(*models.Team)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureSave(
	id int64,
	name string,
	description string,
	slack string,
) {
	arg := &models.Team{
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	fake := &models.Team{
		ID:          id,
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	m.On(`Save`, arg).Return(fake, nil)
}

func (m *RepositoryMock) Update(_ context.Context, team *models.Team) (*models.Team, error) {
	args := m.Called(team)
	res, _ := args.Get(0).(*models.Team)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureUpdate(
	id int64,
	name string,
	description string,
	slack string,
	created time.Time,
	updated time.Time,
) {
	arg := &models.Team{
		ID:          id,
		Name:        name,
		Description: description,
		Slack:       slack,
	}

	fake := &models.Team{
		ID:          id,
		Name:        name,
		Description: description,
		Slack:       slack,
		Created:     created,
		Updated:     updated,
	}

	m.On(`Update`, arg).Return(fake, nil)
}

func (m *RepositoryMock) Remove(_ context.Context, teamID int64) (int64, error) {
	args := m.Called(teamID)
	res, _ := args.Get(0).(int64)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureRemove(id int64) {
	m.On(`Remove`, id).Return(id, nil)
}

func (m *RepositoryMock) Get(
	_ context.Context,
	filter *v1.TeamFilter,
	limit, offset uint,
	sort, order string,
) ([]models.Team, error) {
	args := m.Called(filter, limit, offset, sort, order)
	res, _ := args.Get(0).([]models.Team)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureGet(
	id int64,
	name string,
	description string,
	slack string,
	created time.Time,
	updated time.Time,
) {
	arg := &v1.TeamFilter{
		Ids: []int64{id},
	}

	fake := []models.Team{
		{
			ID:          id,
			Name:        name,
			Description: description,
			Slack:       slack,
			Created:     created,
			Updated:     updated,
		},
	}

	m.On(`Get`, arg, uint(0), uint(0), "", "").Return(fake, nil)
}
