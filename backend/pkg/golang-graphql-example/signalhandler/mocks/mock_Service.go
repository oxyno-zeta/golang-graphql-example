// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler Service
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	os "os"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
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

// ActiveRequestCounterMiddleware mocks base method.
func (m *MockService) ActiveRequestCounterMiddleware(ignoredPathList []string) gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveRequestCounterMiddleware", ignoredPathList)
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// ActiveRequestCounterMiddleware indicates an expected call of ActiveRequestCounterMiddleware.
func (mr *MockServiceMockRecorder) ActiveRequestCounterMiddleware(ignoredPathList any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveRequestCounterMiddleware", reflect.TypeOf((*MockService)(nil).ActiveRequestCounterMiddleware), ignoredPathList)
}

// DecreaseActiveRequestCounter mocks base method.
func (m *MockService) DecreaseActiveRequestCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DecreaseActiveRequestCounter")
}

// DecreaseActiveRequestCounter indicates an expected call of DecreaseActiveRequestCounter.
func (mr *MockServiceMockRecorder) DecreaseActiveRequestCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseActiveRequestCounter", reflect.TypeOf((*MockService)(nil).DecreaseActiveRequestCounter))
}

// GetStoppingSystemContext mocks base method.
func (m *MockService) GetStoppingSystemContext() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoppingSystemContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// GetStoppingSystemContext indicates an expected call of GetStoppingSystemContext.
func (mr *MockServiceMockRecorder) GetStoppingSystemContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoppingSystemContext", reflect.TypeOf((*MockService)(nil).GetStoppingSystemContext))
}

// IncreaseActiveRequestCounter mocks base method.
func (m *MockService) IncreaseActiveRequestCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncreaseActiveRequestCounter")
}

// IncreaseActiveRequestCounter indicates an expected call of IncreaseActiveRequestCounter.
func (mr *MockServiceMockRecorder) IncreaseActiveRequestCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseActiveRequestCounter", reflect.TypeOf((*MockService)(nil).IncreaseActiveRequestCounter))
}

// InitializeOnce mocks base method.
func (m *MockService) InitializeOnce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeOnce")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeOnce indicates an expected call of InitializeOnce.
func (mr *MockServiceMockRecorder) InitializeOnce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeOnce", reflect.TypeOf((*MockService)(nil).InitializeOnce))
}

// IsStoppingSystem mocks base method.
func (m *MockService) IsStoppingSystem() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStoppingSystem")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStoppingSystem indicates an expected call of IsStoppingSystem.
func (mr *MockServiceMockRecorder) IsStoppingSystem() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStoppingSystem", reflect.TypeOf((*MockService)(nil).IsStoppingSystem))
}

// OnExit mocks base method.
func (m *MockService) OnExit(hook func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnExit", hook)
}

// OnExit indicates an expected call of OnExit.
func (mr *MockServiceMockRecorder) OnExit(hook any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnExit", reflect.TypeOf((*MockService)(nil).OnExit), hook)
}

// OnSignal mocks base method.
func (m *MockService) OnSignal(signal os.Signal, hook func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSignal", signal, hook)
}

// OnSignal indicates an expected call of OnSignal.
func (mr *MockServiceMockRecorder) OnSignal(signal, hook any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSignal", reflect.TypeOf((*MockService)(nil).OnSignal), signal, hook)
}
