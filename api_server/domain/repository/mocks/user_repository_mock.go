// Code generated by MockGen. DO NOT EDIT.
// Source: api_server/domain/repository/user_repository.go
//
// Generated by this command:
//
//	mockgen -source=api_server/domain/repository/user_repository.go -destination=api_server/domain/repository/mocks/user_repository_mock.go
//
// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "pinnacle-play/domain/model"
	reflect "reflect"
	"github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockUserRepository) Save(ctx context.Context, name string, groupId model.GroupID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, name, groupId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUserRepositoryMockRecorder) Save(ctx, name, groupId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), ctx, name, groupId)
}
