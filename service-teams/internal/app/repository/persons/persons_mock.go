package persons

import (
	"context"
	"time"

	"github.com/bgoldovsky/dutyer/service-teams/internal/app/models"
	v1 "github.com/bgoldovsky/dutyer/service-teams/internal/generated/rpc/v1"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func NewMock() *RepositoryMock {
	return &RepositoryMock{}
}

func (m *RepositoryMock) Save(_ context.Context, person *models.Person) (*models.Person, error) {
	args := m.Called(person)
	res, _ := args.Get(0).(*models.Person)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureSave(
	id int64,
	firstName string,
	middleName *string,
	lastName string,
	birthday time.Time,
	email *string,
	phone *string,
	slack string,
	role models.Role,
	teamID *int64,
) {
	arg := &models.Person{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Birthday:   &birthday,
		Email:      email,
		Phone:      phone,
		Slack:      slack,
		Role:       role,
		TeamID:     teamID,
	}

	fake := &models.Person{
		ID:         id,
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Birthday:   &birthday,
		Email:      email,
		Phone:      phone,
		Slack:      slack,
		Role:       role,
		TeamID:     teamID,
	}

	m.On(`Save`, arg).Return(fake, nil)
}

func (m *RepositoryMock) Update(_ context.Context, person *models.Person) (*models.Person, error) {
	args := m.Called(person)
	res, _ := args.Get(0).(*models.Person)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureUpdate(
	id int64,
	firstName string,
	middleName *string,
	lastName string,
	birthday time.Time,
	email *string,
	phone *string,
	slack string,
	role models.Role,
	teamID *int64,
	dutyOrder int64,
	isActive bool,
	created time.Time,
	updated time.Time,
) {
	arg := &models.Person{
		ID:         id,
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Birthday:   &birthday,
		Email:      email,
		Phone:      phone,
		Slack:      slack,
		Role:       role,
		TeamID:     teamID,
		DutyOrder:  dutyOrder,
		IsActive:   isActive,
	}

	fake := &models.Person{
		ID:         id,
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Birthday:   &birthday,
		Email:      email,
		Phone:      phone,
		Slack:      slack,
		Role:       role,
		TeamID:     teamID,
		DutyOrder:  dutyOrder,
		IsActive:   isActive,
		Created:    created,
		Updated:    updated,
	}

	m.On(`Update`, arg).Return(fake, nil)
}

func (m *RepositoryMock) Remove(_ context.Context, personID int64) (int64, error) {
	args := m.Called(personID)
	res, _ := args.Get(0).(int64)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureRemove(id int64) {
	m.On(`Remove`, id).Return(id, nil)
}

func (m *RepositoryMock) Get(
	_ context.Context,
	filter *v1.PersonFilter,
	limit, offset uint,
	sort, order string,
) ([]models.Person, error) {
	args := m.Called(filter, limit, offset, sort, order)
	res, _ := args.Get(0).([]models.Person)
	return res, args.Error(1)
}

func (m *RepositoryMock) ConfigureGet(
	id int64,
	firstName string,
	middleName *string,
	lastName string,
	birthday time.Time,
	email *string,
	phone *string,
	slack string,
	role models.Role,
	teamID *int64,
	dutyOrder int64,
	isActive bool,
	created time.Time,
	updated time.Time,
) {
	arg := &v1.PersonFilter{
		PersonIds: []int64{id},
		TeamIds:   []int64{*teamID},
	}

	fake := []models.Person{
		{
			ID:         id,
			FirstName:  firstName,
			MiddleName: middleName,
			LastName:   lastName,
			Birthday:   &birthday,
			Email:      email,
			Phone:      phone,
			Slack:      slack,
			Role:       role,
			TeamID:     teamID,
			DutyOrder:  dutyOrder,
			IsActive:   isActive,
			Created:    created,
			Updated:    updated,
		},
	}

	m.On(`Get`, arg, uint(0), uint(0), "", "").Return(fake, nil)
}
