package persons

import (
	"context"
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	mockPersons "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/persons/persons_mock"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

var person = models.Person{
	ID:        777,
	TeamID:    1111,
	FirstName: "Boris",
	LastName:  "Gold",
	Slack:     "QWERTY",
	Role:      2,
	DutyOrder: 15,
	IsActive:  true,
	Created:   time.Now(),
	Updated:   time.Now(),
}

func TestService_GetPersons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	repoMock := mockPersons.NewMockRepository(ctrl)
	repoMock.EXPECT().
		GetList(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]models.Person{person}, nil)

	service := New(repoMock)

	created, _ := ptypes.TimestampProto(person.Created)
	updated, _ := ptypes.TimestampProto(person.Updated)

	request := &v1.GetPersonsRequest{
		Filter: &v1.PersonFilter{
			PersonIds: []int64{person.ID},
		},
	}

	act, err := service.GetPersons(context.Background(), request)

	assert.NoError(t, err)
	assert.NotEmpty(t, act.Persons)
	assert.Equal(t, person.ID, act.Persons[0].Id)
	assert.Equal(t, person.TeamID, act.Persons[0].TeamId)
	assert.Equal(t, person.FirstName, act.Persons[0].FirstName)
	assert.Equal(t, person.LastName, act.Persons[0].LastName)
	assert.Equal(t, person.Slack, act.Persons[0].Slack)
	assert.Equal(t, int64(person.Role), int64(act.Persons[0].Role))
	assert.Equal(t, person.IsActive, act.Persons[0].IsActive)
	assert.Equal(t, created, act.Persons[0].Created)
	assert.Equal(t, updated, act.Persons[0].Updated)
}

func TestService_AddPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	repoMock := mockPersons.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Save(ctx, gomock.Any()).
		Return(&person, nil)

	service := New(repoMock)

	request := &v1.AddPersonRequest{
		FirstName: person.LastName,
		LastName:  person.LastName,
		Slack:     person.Slack,
		Role:      v1.Role(person.Role),
		IsActive:  person.IsActive,
	}

	act, err := service.AddPerson(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, person.ID, act.Id)
}

func TestService_UpdatePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	repoMock := mockPersons.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Update(ctx, gomock.Any()).
		Return(&person, nil)

	service := New(repoMock)

	request := &v1.UpdatePersonRequest{
		Id:        person.ID,
		FirstName: person.LastName,
		LastName:  person.LastName,
		Slack:     person.Slack,
		Role:      v1.Role(person.Role),
		IsActive:  person.IsActive,
	}

	act, err := service.UpdatePerson(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, act)
}

func TestService_RemovePerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	repoMock := mockPersons.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Remove(ctx, person.ID).
		Return(person.ID, nil)

	service := New(repoMock)

	request := &v1.RemovePersonRequest{Id: person.ID}

	act, err := service.RemoverPerson(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, act)
}

func TestService_GetPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	repoMock := mockPersons.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Get(ctx, person.ID).
		Return(&person, nil)

	service := New(repoMock)

	created, _ := ptypes.TimestampProto(person.Created)
	updated, _ := ptypes.TimestampProto(person.Updated)

	request := &v1.GetPersonRequest{Id: person.ID}

	act, err := service.GetPerson(context.Background(), request)

	assert.NoError(t, err)
	assert.NotEmpty(t, act.Person)
	assert.Equal(t, person.ID, act.Person.Id)
	assert.Equal(t, person.TeamID, act.Person.TeamId)
	assert.Equal(t, person.FirstName, act.Person.FirstName)
	assert.Equal(t, person.LastName, act.Person.LastName)
	assert.Equal(t, person.Slack, act.Person.Slack)
	assert.Equal(t, int64(person.Role), int64(act.Person.Role))
	assert.Equal(t, person.IsActive, act.Person.IsActive)
	assert.Equal(t, created, act.Person.Created)
	assert.Equal(t, updated, act.Person.Updated)
}
