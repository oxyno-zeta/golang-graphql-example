// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing (interfaces: Trace)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tracing "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
)

// MockTrace is a mock of Trace interface.
type MockTrace struct {
	ctrl     *gomock.Controller
	recorder *MockTraceMockRecorder
}

// MockTraceMockRecorder is the mock recorder for MockTrace.
type MockTraceMockRecorder struct {
	mock *MockTrace
}

// NewMockTrace creates a new mock instance.
func NewMockTrace(ctrl *gomock.Controller) *MockTrace {
	mock := &MockTrace{ctrl: ctrl}
	mock.recorder = &MockTraceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrace) EXPECT() *MockTraceMockRecorder {
	return m.recorder
}

// Finish mocks base method.
func (m *MockTrace) Finish() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finish")
}

// Finish indicates an expected call of Finish.
func (mr *MockTraceMockRecorder) Finish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockTrace)(nil).Finish))
}

// GetChildTrace mocks base method.
func (m *MockTrace) GetChildTrace(arg0 string) tracing.Trace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildTrace", arg0)
	ret0, _ := ret[0].(tracing.Trace)
	return ret0
}

// GetChildTrace indicates an expected call of GetChildTrace.
func (mr *MockTraceMockRecorder) GetChildTrace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildTrace", reflect.TypeOf((*MockTrace)(nil).GetChildTrace), arg0)
}

// GetTraceID mocks base method.
func (m *MockTrace) GetTraceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTraceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTraceID indicates an expected call of GetTraceID.
func (mr *MockTraceMockRecorder) GetTraceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraceID", reflect.TypeOf((*MockTrace)(nil).GetTraceID))
}

// InjectInHTTPHeader mocks base method.
func (m *MockTrace) InjectInHTTPHeader(arg0 http.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectInHTTPHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectInHTTPHeader indicates an expected call of InjectInHTTPHeader.
func (mr *MockTraceMockRecorder) InjectInHTTPHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectInHTTPHeader", reflect.TypeOf((*MockTrace)(nil).InjectInHTTPHeader), arg0)
}

// InjectInTextMap mocks base method.
func (m *MockTrace) InjectInTextMap(arg0 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectInTextMap", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectInTextMap indicates an expected call of InjectInTextMap.
func (mr *MockTraceMockRecorder) InjectInTextMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectInTextMap", reflect.TypeOf((*MockTrace)(nil).InjectInTextMap), arg0)
}

// MarkAsError mocks base method.
func (m *MockTrace) MarkAsError() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkAsError")
}

// MarkAsError indicates an expected call of MarkAsError.
func (mr *MockTraceMockRecorder) MarkAsError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsError", reflect.TypeOf((*MockTrace)(nil).MarkAsError))
}

// SetTag mocks base method.
func (m *MockTrace) SetTag(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTag", arg0, arg1)
}

// SetTag indicates an expected call of SetTag.
func (mr *MockTraceMockRecorder) SetTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTag", reflect.TypeOf((*MockTrace)(nil).SetTag), arg0, arg1)
}

// SetTags mocks base method.
func (m *MockTrace) SetTags(arg0 map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTags", arg0)
}

// SetTags indicates an expected call of SetTags.
func (mr *MockTraceMockRecorder) SetTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTags", reflect.TypeOf((*MockTrace)(nil).SetTags), arg0)
}
