package duties

import (
	"context"
	"errors"
	"testing"

	mockPublisher "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/publisher/publisher_mock"
	mockDuties "github.com/bgoldovsky/dutyer/service-dutyer/internal/app/repository/duties/duties_mock"
	v1 "github.com/bgoldovsky/dutyer/service-dutyer/internal/generated/rpc/v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

	// TODO: настроить моки отправки event
	pubMock := mockPublisher.NewMockPublisher(ctrl)

	service := New(nil, repoMock, nil, pubMock)

	err := service.Swap(ctx, req)
	assert.NoError(t, err)
}

// TODO: Дописать тесты сервиса
func TestService_Assign_Error(t *testing.T) {

}

// TODO: Дописать тесты сервиса
func TestService_Assign_Success(t *testing.T) {

}

// TODO: Дописать тесты сервиса
func TestService_AssignNextDuties_Error(t *testing.T) {

}

// TODO: Дописать тесты сервиса
func TestService_AssignNextDuties_Success(t *testing.T) {

}

// TODO: Дописать тесты сервиса
func TestService_GetDuties_Error(t *testing.T) {

}

// TODO: Дописать тесты сервиса
func TestService_GetDuties_Success(t *testing.T) {

}

// TODO: Дописать тесты сервиса
func TestService_GetCurrentDuty_Error(t *testing.T) {

}

// TODO: Дописать тесты сервиса
func TestService_GetCurrentDuty_Success(t *testing.T) {

}
