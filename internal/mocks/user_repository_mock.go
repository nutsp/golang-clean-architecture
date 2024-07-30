// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/user_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nutsp/golang-clean-architecture/internal/models"
	repositories "github.com/nutsp/golang-clean-architecture/internal/repositories"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Atomic mocks base method.
func (m *MockIUserRepository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(repositories.IUserRepository) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Atomic", ctx, opt, repo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Atomic indicates an expected call of Atomic.
func (mr *MockIUserRepositoryMockRecorder) Atomic(ctx, opt, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Atomic", reflect.TypeOf((*MockIUserRepository)(nil).Atomic), ctx, opt, repo)
}

// GetAll mocks base method.
func (m *MockIUserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIUserRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIUserRepository)(nil).GetAll), ctx)
}

// GetByID mocks base method.
func (m *MockIUserRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIUserRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIUserRepository)(nil).GetByID), ctx, id)
}

// Save mocks base method.
func (m *MockIUserRepository) Save(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIUserRepositoryMockRecorder) Save(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIUserRepository)(nil).Save), ctx, user)
}

// UpdateByID mocks base method.
func (m *MockIUserRepository) UpdateByID(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockIUserRepositoryMockRecorder) UpdateByID(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockIUserRepository)(nil).UpdateByID), ctx, user)
}
