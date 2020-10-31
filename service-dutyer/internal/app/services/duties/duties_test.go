package duties

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
	mockPublisher "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/publisher/publisher_mock"
	mockDuties "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/duties/duties_mock"
	mockPersons "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/persons/persons_mock"
	mockTeams "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/teams/teams_mock"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	duty = models.Duty{
		TeamID:    111,
		PersonID:  222,
		FirstName: "John",
		LastName:  "Doe",
		Slack:     "QWERTY",
		Order:     5,
		Month:     12,
		Day:       31,
		Created:   time.Now(),
		Updated:   time.Now(),
	}

	person = models.Person{
		ID:        222,
		TeamID:    111,
		FirstName: "John",
		LastName:  "Doe",
		Slack:     "QWERTY",
		Role:      4,
		DutyOrder: 5,
		IsActive:  true,
		Created:   time.Now(),
		Updated:   time.Now(),
	}

	team = models.Team{
		ID:          111,
		Name:        "Dream Team",
		Description: "My dream team",
		Slack:       "XXXYYYZZZ",
		Created:     time.Now(),
		Updated:     time.Now(),
	}
)

func TestService_Swap_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	req := &v1.SwapRequest{
		TeamId:         111,
		FirstPersonId:  222,
		SecondPersonId: 333,
	}

	expErr := errors.New("my repo error")

	repoMock := mockDuties.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Swap(ctx, req.TeamId, req.FirstPersonId, req.SecondPersonId).
		Return(expErr)

	pubMock := mockPublisher.NewMockPublisher(ctrl)

	service := New(nil, repoMock, nil, pubMock)

	err := service.Swap(ctx, req)
	assert.Error(t, err)
	assert.EqualError(t, err, "swap persons error: my repo error")
}

func TestService_Swap_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	req := &v1.SwapRequest{
		TeamId:         111,
		FirstPersonId:  222,
		SecondPersonId: 333,
	}

	repoMock := mockDuties.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Swap(ctx, req.TeamId, req.FirstPersonId, req.SecondPersonId).
		Return(nil)

	pubMock := mockPublisher.NewMockPublisher(ctrl)

	service := New(nil, repoMock, nil, pubMock)

	err := service.Swap(ctx, req)
	assert.NoError(t, err)
}

func TestService_Assign_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	expErr := errors.New("my repo error")

	personRepoMock := mockPersons.NewMockRepository(ctrl)
	personRepoMock.EXPECT().
		Get(ctx, person.ID).
		Return(&person, nil)

	dutyRepoMock := mockDuties.NewMockRepository(ctrl)
	dutyRepoMock.EXPECT().
		Save(ctx, gomock.Any()).
		Return(nil)

	pubMock := mockPublisher.NewMockPublisher(ctrl)
	pubMock.EXPECT().
		Publish(gomock.Any(), person.TeamID, gomock.Any()).
		Return(expErr)

	service := New(personRepoMock, dutyRepoMock, nil, pubMock)

	err := service.Assign(ctx, person.TeamID, person.ID, time.Now())
	assert.EqualError(t, err, "publish duty.assigned to duties error: my repo error")
}

func TestService_Assign_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	personRepoMock := mockPersons.NewMockRepository(ctrl)
	personRepoMock.EXPECT().
		Get(ctx, person.ID).
		Return(&person, nil)

	dutyRepoMock := mockDuties.NewMockRepository(ctrl)
	dutyRepoMock.EXPECT().
		Save(ctx, gomock.Any()).
		Return(nil)

	pubMock := mockPublisher.NewMockPublisher(ctrl)
	pubMock.EXPECT().
		Publish(gomock.Any(), person.TeamID, gomock.Any()).
		Return(nil)

	service := New(personRepoMock, dutyRepoMock, nil, pubMock)

	err := service.Assign(ctx, person.TeamID, person.ID, time.Now())
	assert.NoError(t, err)
}

func TestService_AssignNextDuties_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	expErr := errors.New("my repo error")

	personRepoMock := mockPersons.NewMockRepository(ctrl)
	dutyRepoMock := mockDuties.NewMockRepository(ctrl)
	pubMock := mockPublisher.NewMockPublisher(ctrl)

	teamRepoMock := mockTeams.NewMockRepository(ctrl)
	teamRepoMock.EXPECT().
		GetList(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, expErr)

	service := New(personRepoMock, dutyRepoMock, teamRepoMock, pubMock)

	err := service.AssignNextDuties(ctx, time.Now())
	assert.EqualError(t, err, "get team list error: my repo error")
}

func TestService_AssignNextDuties_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	personRepoMock := mockPersons.NewMockRepository(ctrl)
	personRepoMock.EXPECT().
		GetList(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]models.Person{person}, nil)

	dutyRepoMock := mockDuties.NewMockRepository(ctrl)
	dutyRepoMock.EXPECT().
		Get(ctx, gomock.Any()).
		Return(&duty, nil)
	dutyRepoMock.EXPECT().
		Save(ctx, gomock.Any()).
		Return(nil)

	teamRepoMock := mockTeams.NewMockRepository(ctrl)
	teamRepoMock.EXPECT().
		GetList(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]models.Team{team}, nil)

	pubMock := mockPublisher.NewMockPublisher(ctrl)
	pubMock.EXPECT().
		Publish(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	service := New(personRepoMock, dutyRepoMock, teamRepoMock, pubMock)

	err := service.AssignNextDuties(ctx, time.Now())
	assert.NoError(t, err)
}

func TestService_GetDuties_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	expErr := errors.New("my repo error")

	req := &v1.GetCurrentDutyRequest{
		TeamId: 111,
	}

	personRepoMock := mockPersons.NewMockRepository(ctrl)
	personRepoMock.EXPECT().
		GetList(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, expErr)

	dutyRepoMock := mockDuties.NewMockRepository(ctrl)
	dutyRepoMock.EXPECT().
		Get(ctx, req.TeamId).
		Return(&duty, nil)

	pubMock := mockPublisher.NewMockPublisher(ctrl)

	service := New(personRepoMock, dutyRepoMock, nil, pubMock)

	act, err := service.GetDuties(ctx, duty.TeamID, 1, duty.Created)
	assert.EqualError(t, err, "get persons error: my repo error")
	assert.Nil(t, act)
}

func TestService_GetDuties_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	req := &v1.GetCurrentDutyRequest{
		TeamId: 111,
	}

	personRepoMock := mockPersons.NewMockRepository(ctrl)
	personRepoMock.EXPECT().
		GetList(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]models.Person{person}, nil)

	dutyRepoMock := mockDuties.NewMockRepository(ctrl)
	dutyRepoMock.EXPECT().
		Get(ctx, req.TeamId).
		Return(&duty, nil)

	pubMock := mockPublisher.NewMockPublisher(ctrl)

	service := New(personRepoMock, dutyRepoMock, nil, pubMock)

	act, err := service.GetDuties(ctx, duty.TeamID, 1, duty.Created)
	assert.NoError(t, err)
	assert.Equal(t, duty.PersonID, act[0].PersonID)
	assert.Equal(t, duty.TeamID, act[0].TeamID)
	assert.Equal(t, duty.Order, act[0].Order)
	assert.Equal(t, duty.Channel, act[0].Channel)
	assert.Equal(t, duty.Slack, act[0].Slack)
}

func TestService_GetCurrentDuty_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	expErr := errors.New("my repo error")

	req := &v1.GetCurrentDutyRequest{
		TeamId: 111,
	}

	repoMock := mockDuties.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Get(ctx, req.TeamId).
		Return(nil, expErr)

	pubMock := mockPublisher.NewMockPublisher(ctrl)

	service := New(nil, repoMock, nil, pubMock)

	act, err := service.GetCurrentDuty(ctx, duty.TeamID)
	assert.EqualError(t, err, "get current duty error: my repo error")
	assert.Nil(t, act)
}

func TestService_GetCurrentDuty_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	req := &v1.GetCurrentDutyRequest{
		TeamId: 111,
	}

	repoMock := mockDuties.NewMockRepository(ctrl)
	repoMock.EXPECT().
		Get(ctx, req.TeamId).
		Return(&duty, nil)

	pubMock := mockPublisher.NewMockPublisher(ctrl)

	service := New(nil, repoMock, nil, pubMock)

	act, err := service.GetCurrentDuty(ctx, duty.TeamID)
	assert.NoError(t, err)
	assert.Equal(t, &duty, act)
}
