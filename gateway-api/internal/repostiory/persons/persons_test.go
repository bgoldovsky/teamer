package persons

import (
	"time"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func NewMock(
	id int64,
	name string,
	description string,
	slack string,
	created time.Time,
	updated time.Time,
) *RepositoryMock {
	m := &RepositoryMock{}
	view := []models.TeamView{
		{
			ID:          id,
			Name:        name,
			Description: description,
			Slack:       slack,
			Crated:      created,
			Updated:     updated,
		},
	}

	m.On(`Clear`).Return(nil)
	m.On(`Save`, view).Return(nil)
	m.On(`Get`).Return(view, nil)

	return m
}

func (m *RepositoryMock) Save(teams []models.TeamView) error {
	args := m.Called(teams)
	return args.Error(0)
}

func (m *RepositoryMock) Get() ([]models.TeamView, error) {
	args := m.Called()
	res, _ := args.Get(0).([]models.TeamView)
	return res, args.Error(1)
}

func (m *RepositoryMock) Clear() error {
	args := m.Called()
	return args.Error(0)
}
