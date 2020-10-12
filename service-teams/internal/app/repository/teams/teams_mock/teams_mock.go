// Code generated by MockGen. DO NOT EDIT.
// Source: teams.go

// Package mock_teams is a generated GoMock package.
package mock_teams

import (
	context "context"
	models "github.com/bgoldovsky/dutyer/service-teams/internal/app/models"
	v1 "github.com/bgoldovsky/dutyer/service-teams/internal/generated/rpc/v1"
	gomock "github.com/golang/mock/gomock"
	v4 "github.com/jackc/pgx/v4"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(ctx context.Context, teamID int64) (*models.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, teamID)
	ret0, _ := ret[0].(*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(ctx, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), ctx, teamID)
}

// GetList mocks base method
func (m *MockRepository) GetList(ctx context.Context, filter *v1.TeamFilter, limit, offset uint, sort, order string) ([]models.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, filter, limit, offset, sort, order)
	ret0, _ := ret[0].([]models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList
func (mr *MockRepositoryMockRecorder) GetList(ctx, filter, limit, offset, sort, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockRepository)(nil).GetList), ctx, filter, limit, offset, sort, order)
}

// Save mocks base method
func (m *MockRepository) Save(ctx context.Context, team *models.Team) (*models.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, team)
	ret0, _ := ret[0].(*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockRepositoryMockRecorder) Save(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), ctx, team)
}

// Update mocks base method
func (m *MockRepository) Update(ctx context.Context, team *models.Team) (*models.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, team)
	ret0, _ := ret[0].(*models.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, team)
}

// Remove mocks base method
func (m *MockRepository) Remove(ctx context.Context, teamID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, teamID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove
func (mr *MockRepositoryMockRecorder) Remove(ctx, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockRepository)(nil).Remove), ctx, teamID)
}

// Mockqueryer is a mock of queryer interface
type Mockqueryer struct {
	ctrl     *gomock.Controller
	recorder *MockqueryerMockRecorder
}

// MockqueryerMockRecorder is the mock recorder for Mockqueryer
type MockqueryerMockRecorder struct {
	mock *Mockqueryer
}

// NewMockqueryer creates a new mock instance
func NewMockqueryer(ctrl *gomock.Controller) *Mockqueryer {
	mock := &Mockqueryer{ctrl: ctrl}
	mock.recorder = &MockqueryerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockqueryer) EXPECT() *MockqueryerMockRecorder {
	return m.recorder
}

// Query mocks base method
func (m *Mockqueryer) Query(ctx context.Context, sql string, args ...interface{}) (v4.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(v4.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockqueryerMockRecorder) Query(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*Mockqueryer)(nil).Query), varargs...)
}

// QueryRow mocks base method
func (m *Mockqueryer) QueryRow(ctx context.Context, sql string, args ...interface{}) v4.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(v4.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow
func (mr *MockqueryerMockRecorder) QueryRow(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*Mockqueryer)(nil).QueryRow), varargs...)
}
