// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-playbooks/server/app (interfaces: UserInfoService)

// Package mock_app is a generated GoMock package.
package mock_app

import (
	gomock "github.com/golang/mock/gomock"
	app "github.com/mattermost/mattermost-plugin-playbooks/server/app"
	reflect "reflect"
)

// MockUserInfoService is a mock of UserInfoService interface
type MockUserInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockUserInfoServiceMockRecorder
}

// MockUserInfoServiceMockRecorder is the mock recorder for MockUserInfoService
type MockUserInfoServiceMockRecorder struct {
	mock *MockUserInfoService
}

// NewMockUserInfoService creates a new mock instance
func NewMockUserInfoService(ctrl *gomock.Controller) *MockUserInfoService {
	mock := &MockUserInfoService{ctrl: ctrl}
	mock.recorder = &MockUserInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserInfoService) EXPECT() *MockUserInfoServiceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUserInfoService) Get(arg0 string) (app.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(app.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUserInfoServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserInfoService)(nil).Get), arg0)
}

// Upsert mocks base method
func (m *MockUserInfoService) Upsert(arg0 app.UserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockUserInfoServiceMockRecorder) Upsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockUserInfoService)(nil).Upsert), arg0)
}