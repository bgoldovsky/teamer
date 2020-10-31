// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_duties is a generated GoMock package.
package duties

import (
	context "context"
	reflect "reflect"

	models "github.com/bgoldovsky/dutyer/service-slack-bot/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetCurrentDuty mocks base method
func (m *MockClient) GetCurrentDuty(ctx context.Context, teamID int64) (*models.Duty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentDuty", ctx, teamID)
	ret0, _ := ret[0].(*models.Duty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentDuty indicates an expected call of GetCurrentDuty
func (mr *MockClientMockRecorder) GetCurrentDuty(ctx, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentDuty", reflect.TypeOf((*MockClient)(nil).GetCurrentDuty), ctx, teamID)
}