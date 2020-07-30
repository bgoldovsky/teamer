package publisher

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func NewMock(eventName string, entityID int64, topic string) *Mock {
	m := &Mock{}
	m.On(`Publish`, eventName, entityID, topic).Return(nil)
	return m
}

func (m *Mock) Publish(eventName string, entityID int64, topic string) error {
	args := m.Called(eventName, entityID, topic)
	return args.Error(0)
}
