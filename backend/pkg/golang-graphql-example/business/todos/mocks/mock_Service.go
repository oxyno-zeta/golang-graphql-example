// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	todos "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos"
	models "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	pagination "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database/pagination"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
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

// Close mocks base method.
func (m *MockService) Close(arg0 context.Context, arg1 string) (*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0, arg1)
	ret0, _ := ret[0].(*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close.
func (mr *MockServiceMockRecorder) Close(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockService)(nil).Close), arg0, arg1)
}

// Create mocks base method.
func (m *MockService) Create(arg0 context.Context, arg1 *todos.InputCreateTodo) (*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), arg0, arg1)
}

// Find mocks base method.
func (m *MockService) Find(arg0 context.Context, arg1 []*models.SortOrder, arg2 *models.Filter, arg3 *models.Projection) ([]*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockServiceMockRecorder) Find(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockService)(nil).Find), arg0, arg1, arg2, arg3)
}

// FindByID mocks base method.
func (m *MockService) FindByID(arg0 context.Context, arg1 string, arg2 *models.Projection) (*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockServiceMockRecorder) FindByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockService)(nil).FindByID), arg0, arg1, arg2)
}

// GetAllPaginated mocks base method.
func (m *MockService) GetAllPaginated(arg0 context.Context, arg1 *pagination.PageInput, arg2 []*models.SortOrder, arg3 *models.Filter, arg4 *models.Projection) ([]*models.Todo, *pagination.PageOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPaginated", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*models.Todo)
	ret1, _ := ret[1].(*pagination.PageOutput)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllPaginated indicates an expected call of GetAllPaginated.
func (mr *MockServiceMockRecorder) GetAllPaginated(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPaginated", reflect.TypeOf((*MockService)(nil).GetAllPaginated), arg0, arg1, arg2, arg3, arg4)
}

// Update mocks base method.
func (m *MockService) Update(arg0 context.Context, arg1 *todos.InputUpdateTodo) (*models.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*models.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), arg0, arg1)
}
