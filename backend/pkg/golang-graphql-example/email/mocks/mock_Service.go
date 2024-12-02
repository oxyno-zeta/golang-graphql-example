// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email Service
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	email "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockService) Check() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check")
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockServiceMockRecorder) Check() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockService)(nil).Check))
}

// InitializeAndReload mocks base method.
func (m *MockService) InitializeAndReload() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeAndReload")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeAndReload indicates an expected call of InitializeAndReload.
func (mr *MockServiceMockRecorder) InitializeAndReload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeAndReload", reflect.TypeOf((*MockService)(nil).InitializeAndReload))
}

// NewEmail mocks base method.
func (m *MockService) NewEmail() email.Email {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEmail")
	ret0, _ := ret[0].(email.Email)
	return ret0
}

// NewEmail indicates an expected call of NewEmail.
func (mr *MockServiceMockRecorder) NewEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEmail", reflect.TypeOf((*MockService)(nil).NewEmail))
}

// Send mocks base method.
func (m *MockService) Send(em email.Email) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", em)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockServiceMockRecorder) Send(em any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockService)(nil).Send), em)
}
