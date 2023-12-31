// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor/sql (interfaces: Lock)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_Lock.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/lockdistributor/sql Lock
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockLock is a mock of Lock interface.
type MockLock struct {
	ctrl     *gomock.Controller
	recorder *MockLockMockRecorder
}

// MockLockMockRecorder is the mock recorder for MockLock.
type MockLockMockRecorder struct {
	mock *MockLock
}

// NewMockLock creates a new mock instance.
func NewMockLock(ctrl *gomock.Controller) *MockLock {
	mock := &MockLock{ctrl: ctrl}
	mock.recorder = &MockLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLock) EXPECT() *MockLockMockRecorder {
	return m.recorder
}

// Acquire mocks base method.
func (m *MockLock) Acquire() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Acquire")
	ret0, _ := ret[0].(error)
	return ret0
}

// Acquire indicates an expected call of Acquire.
func (mr *MockLockMockRecorder) Acquire() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acquire", reflect.TypeOf((*MockLock)(nil).Acquire))
}

// AcquireWithContext mocks base method.
func (m *MockLock) AcquireWithContext(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireWithContext", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcquireWithContext indicates an expected call of AcquireWithContext.
func (mr *MockLockMockRecorder) AcquireWithContext(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireWithContext", reflect.TypeOf((*MockLock)(nil).AcquireWithContext), arg0)
}

// IsAlreadyTaken mocks base method.
func (m *MockLock) IsAlreadyTaken() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlreadyTaken")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAlreadyTaken indicates an expected call of IsAlreadyTaken.
func (mr *MockLockMockRecorder) IsAlreadyTaken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlreadyTaken", reflect.TypeOf((*MockLock)(nil).IsAlreadyTaken))
}

// IsReleased mocks base method.
func (m *MockLock) IsReleased() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReleased")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsReleased indicates an expected call of IsReleased.
func (mr *MockLockMockRecorder) IsReleased() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReleased", reflect.TypeOf((*MockLock)(nil).IsReleased))
}

// Release mocks base method.
func (m *MockLock) Release() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release")
	ret0, _ := ret[0].(error)
	return ret0
}

// Release indicates an expected call of Release.
func (mr *MockLockMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockLock)(nil).Release))
}
