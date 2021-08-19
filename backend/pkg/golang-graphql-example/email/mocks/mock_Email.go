// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email (interfaces: Email)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	email "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/email"
	mail "github.com/xhit/go-simple-mail/v2"
)

// MockEmail is a mock of Email interface.
type MockEmail struct {
	ctrl     *gomock.Controller
	recorder *MockEmailMockRecorder
}

// MockEmailMockRecorder is the mock recorder for MockEmail.
type MockEmailMockRecorder struct {
	mock *MockEmail
}

// NewMockEmail creates a new mock instance.
func NewMockEmail(ctrl *gomock.Controller) *MockEmail {
	mock := &MockEmail{ctrl: ctrl}
	mock.recorder = &MockEmailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmail) EXPECT() *MockEmailMockRecorder {
	return m.recorder
}

// AddAttachment mocks base method.
func (m *MockEmail) AddAttachment(arg0 []byte, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAttachment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAttachment indicates an expected call of AddAttachment.
func (mr *MockEmailMockRecorder) AddAttachment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttachment", reflect.TypeOf((*MockEmail)(nil).AddAttachment), arg0, arg1, arg2)
}

// AddAttachmentFile mocks base method.
func (m *MockEmail) AddAttachmentFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAttachmentFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAttachmentFile indicates an expected call of AddAttachmentFile.
func (mr *MockEmailMockRecorder) AddAttachmentFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttachmentFile", reflect.TypeOf((*MockEmail)(nil).AddAttachmentFile), arg0)
}

// AddBcc mocks base method.
func (m *MockEmail) AddBcc(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddBcc", varargs...)
}

// AddBcc indicates an expected call of AddBcc.
func (mr *MockEmailMockRecorder) AddBcc(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBcc", reflect.TypeOf((*MockEmail)(nil).AddBcc), arg0...)
}

// AddCc mocks base method.
func (m *MockEmail) AddCc(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddCc", varargs...)
}

// AddCc indicates an expected call of AddCc.
func (mr *MockEmailMockRecorder) AddCc(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCc", reflect.TypeOf((*MockEmail)(nil).AddCc), arg0...)
}

// AddInlineAttachment mocks base method.
func (m *MockEmail) AddInlineAttachment(arg0 []byte, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInlineAttachment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddInlineAttachment indicates an expected call of AddInlineAttachment.
func (mr *MockEmailMockRecorder) AddInlineAttachment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInlineAttachment", reflect.TypeOf((*MockEmail)(nil).AddInlineAttachment), arg0, arg1, arg2)
}

// AddInlineAttachmentFile mocks base method.
func (m *MockEmail) AddInlineAttachmentFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInlineAttachmentFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddInlineAttachmentFile indicates an expected call of AddInlineAttachmentFile.
func (mr *MockEmailMockRecorder) AddInlineAttachmentFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInlineAttachmentFile", reflect.TypeOf((*MockEmail)(nil).AddInlineAttachmentFile), arg0)
}

// AddTo mocks base method.
func (m *MockEmail) AddTo(arg0 ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddTo", varargs...)
}

// AddTo indicates an expected call of AddTo.
func (mr *MockEmailMockRecorder) AddTo(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTo", reflect.TypeOf((*MockEmail)(nil).AddTo), arg0...)
}

// GetEmail mocks base method.
func (m *MockEmail) GetEmail() *mail.Email {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail")
	ret0, _ := ret[0].(*mail.Email)
	return ret0
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockEmailMockRecorder) GetEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockEmail)(nil).GetEmail))
}

// SetDate mocks base method.
func (m *MockEmail) SetDate(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDate", arg0)
}

// SetDate indicates an expected call of SetDate.
func (mr *MockEmailMockRecorder) SetDate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDate", reflect.TypeOf((*MockEmail)(nil).SetDate), arg0)
}

// SetFrom mocks base method.
func (m *MockEmail) SetFrom(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFrom", arg0)
}

// SetFrom indicates an expected call of SetFrom.
func (mr *MockEmailMockRecorder) SetFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFrom", reflect.TypeOf((*MockEmail)(nil).SetFrom), arg0)
}

// SetHTMLBody mocks base method.
func (m *MockEmail) SetHTMLBody(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHTMLBody", arg0)
}

// SetHTMLBody indicates an expected call of SetHTMLBody.
func (mr *MockEmailMockRecorder) SetHTMLBody(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHTMLBody", reflect.TypeOf((*MockEmail)(nil).SetHTMLBody), arg0)
}

// SetPriority mocks base method.
func (m *MockEmail) SetPriority(arg0 email.Priority) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPriority", arg0)
}

// SetPriority indicates an expected call of SetPriority.
func (mr *MockEmailMockRecorder) SetPriority(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPriority", reflect.TypeOf((*MockEmail)(nil).SetPriority), arg0)
}

// SetReplyTo mocks base method.
func (m *MockEmail) SetReplyTo(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReplyTo", arg0)
}

// SetReplyTo indicates an expected call of SetReplyTo.
func (mr *MockEmailMockRecorder) SetReplyTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReplyTo", reflect.TypeOf((*MockEmail)(nil).SetReplyTo), arg0)
}

// SetSender mocks base method.
func (m *MockEmail) SetSender(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSender", arg0)
}

// SetSender indicates an expected call of SetSender.
func (mr *MockEmailMockRecorder) SetSender(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSender", reflect.TypeOf((*MockEmail)(nil).SetSender), arg0)
}

// SetSubject mocks base method.
func (m *MockEmail) SetSubject(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubject", arg0)
}

// SetSubject indicates an expected call of SetSubject.
func (mr *MockEmailMockRecorder) SetSubject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubject", reflect.TypeOf((*MockEmail)(nil).SetSubject), arg0)
}

// SetTextBody mocks base method.
func (m *MockEmail) SetTextBody(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTextBody", arg0)
}

// SetTextBody indicates an expected call of SetTextBody.
func (mr *MockEmailMockRecorder) SetTextBody(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTextBody", reflect.TypeOf((*MockEmail)(nil).SetTextBody), arg0)
}