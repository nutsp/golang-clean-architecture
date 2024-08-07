// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/mailer_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMailerRepository is a mock of IMailerRepository interface.
type MockIMailerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIMailerRepositoryMockRecorder
}

// MockIMailerRepositoryMockRecorder is the mock recorder for MockIMailerRepository.
type MockIMailerRepositoryMockRecorder struct {
	mock *MockIMailerRepository
}

// NewMockIMailerRepository creates a new mock instance.
func NewMockIMailerRepository(ctrl *gomock.Controller) *MockIMailerRepository {
	mock := &MockIMailerRepository{ctrl: ctrl}
	mock.recorder = &MockIMailerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMailerRepository) EXPECT() *MockIMailerRepositoryMockRecorder {
	return m.recorder
}

// CheckEmailAvailability mocks base method.
func (m *MockIMailerRepository) CheckEmailAvailability(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailAvailability", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailAvailability indicates an expected call of CheckEmailAvailability.
func (mr *MockIMailerRepositoryMockRecorder) CheckEmailAvailability(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailAvailability", reflect.TypeOf((*MockIMailerRepository)(nil).CheckEmailAvailability), ctx, email)
}
