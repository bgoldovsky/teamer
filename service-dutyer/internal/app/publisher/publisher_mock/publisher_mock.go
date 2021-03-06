// Code generated by MockGen. DO NOT EDIT.
// Source: publisher.go

// Package mock_publisher is a generated GoMock package.
package mock_publisher

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPublisher is a mock of Publisher interface
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockPublisher) Publish(eventName string, entityID int64, topic string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", eventName, entityID, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockPublisherMockRecorder) Publish(eventName, entityID, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPublisher)(nil).Publish), eventName, entityID, topic)
}
