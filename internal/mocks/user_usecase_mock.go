// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/user_usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nutsp/golang-clean-architecture/internal/models"
)

// MockIUserUsecase is a mock of IUserUsecase interface.
type MockIUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUserUsecaseMockRecorder
}

// MockIUserUsecaseMockRecorder is the mock recorder for MockIUserUsecase.
type MockIUserUsecaseMockRecorder struct {
	mock *MockIUserUsecase
}

// NewMockIUserUsecase creates a new mock instance.
func NewMockIUserUsecase(ctrl *gomock.Controller) *MockIUserUsecase {
	mock := &MockIUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserUsecase) EXPECT() *MockIUserUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIUserUsecase) CreateUser(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserUsecaseMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserUsecase)(nil).CreateUser), ctx, user)
}

// GetUserInfo mocks base method.
func (m *MockIUserUsecase) GetUserInfo(ctx context.Context, id uint) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockIUserUsecaseMockRecorder) GetUserInfo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockIUserUsecase)(nil).GetUserInfo), ctx, id)
}

// UpdateUserInfo mocks base method.
func (m *MockIUserUsecase) UpdateUserInfo(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockIUserUsecaseMockRecorder) UpdateUserInfo(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockIUserUsecase)(nil).UpdateUserInfo), ctx, user)
}