// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/authentication.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthentication is a mock of Authentication interface
type MockAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationMockRecorder
}

// MockAuthenticationMockRecorder is the mock recorder for MockAuthentication
type MockAuthenticationMockRecorder struct {
	mock *MockAuthentication
}

// NewMockAuthentication creates a new mock instance
func NewMockAuthentication(ctrl *gomock.Controller) *MockAuthentication {
	mock := &MockAuthentication{ctrl: ctrl}
	mock.recorder = &MockAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthentication) EXPECT() *MockAuthenticationMockRecorder {
	return m.recorder
}

// SignIn mocks base method
func (m *MockAuthentication) SignIn(ctx context.Context, userID, password string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, userID, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignIn indicates an expected call of SignIn
func (mr *MockAuthenticationMockRecorder) SignIn(ctx, userID, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthentication)(nil).SignIn), ctx, userID, password)
}
