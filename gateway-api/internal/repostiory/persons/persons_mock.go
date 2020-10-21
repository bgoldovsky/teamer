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
	teamID int64,
	firstName string,
	lastName string,
	slack string,
	role int64,
	created time.Time,
	updated time.Time,
) *RepositoryMock {
	m := &RepositoryMock{}
	view := []models.PersonView{
		{
			ID:        id,
			TeamId:    teamID,
			FirstName: firstName,
			LastName:  lastName,
			Slack:     slack,
			Role:      role,
			Created:   created,
			Updated:   updated,
		},
	}

	m.On(`Clear`).Return(nil)
	m.On(`Save`, view).Return(nil)
	m.On(`Get`).Return(view, nil)

	return m
}

func (m *RepositoryMock) Save(teams []models.PersonView) error {
	args := m.Called(teams)
	return args.Error(0)
}

func (m *RepositoryMock) Get() ([]models.PersonView, error) {
	args := m.Called()
	res, _ := args.Get(0).([]models.PersonView)
	return res, args.Error(1)
}

func (m *RepositoryMock) Clear() error {
	args := m.Called()
	return args.Error(0)
}
