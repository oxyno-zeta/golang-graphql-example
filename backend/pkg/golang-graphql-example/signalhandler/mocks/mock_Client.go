// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	os "os"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ActiveRequestCounterMiddleware mocks base method.
func (m *MockClient) ActiveRequestCounterMiddleware() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveRequestCounterMiddleware")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// ActiveRequestCounterMiddleware indicates an expected call of ActiveRequestCounterMiddleware.
func (mr *MockClientMockRecorder) ActiveRequestCounterMiddleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveRequestCounterMiddleware", reflect.TypeOf((*MockClient)(nil).ActiveRequestCounterMiddleware))
}

// Initialize mocks base method.
func (m *MockClient) Initialize() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockClientMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockClient)(nil).Initialize))
}

// IsStoppingSystem mocks base method.
func (m *MockClient) IsStoppingSystem() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStoppingSystem")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStoppingSystem indicates an expected call of IsStoppingSystem.
func (mr *MockClientMockRecorder) IsStoppingSystem() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStoppingSystem", reflect.TypeOf((*MockClient)(nil).IsStoppingSystem))
}

// OnExit mocks base method.
func (m *MockClient) OnExit(arg0 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnExit", arg0)
}

// OnExit indicates an expected call of OnExit.
func (mr *MockClientMockRecorder) OnExit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnExit", reflect.TypeOf((*MockClient)(nil).OnExit), arg0)
}

// OnSignal mocks base method.
func (m *MockClient) OnSignal(arg0 os.Signal, arg1 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSignal", arg0, arg1)
}

// OnSignal indicates an expected call of OnSignal.
func (mr *MockClientMockRecorder) OnSignal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSignal", reflect.TypeOf((*MockClient)(nil).OnSignal), arg0, arg1)
}
