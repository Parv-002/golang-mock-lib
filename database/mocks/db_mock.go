// Code generated by MockGen. DO NOT EDIT.
// Source: example.com/m/database (interfaces: Db)

// Package databaseMock is a generated GoMock package.
package databaseMock

import (
	reflect "reflect"

	database "example.com/m/database"
	gomock "github.com/golang/mock/gomock"
)

// MockDb is a mock of Db interface.
type MockDb struct {
	ctrl     *gomock.Controller
	recorder *MockDbMockRecorder
}

// MockDbMockRecorder is the mock recorder for MockDb.
type MockDbMockRecorder struct {
	mock *MockDb
}

// NewMockDb creates a new mock instance.
func NewMockDb(ctrl *gomock.Controller) *MockDb {
	mock := &MockDb{ctrl: ctrl}
	mock.recorder = &MockDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDb) EXPECT() *MockDbMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDb) Get() []database.Student {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]database.Student)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockDbMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDb)(nil).Get))
}

// GetById mocks base method.
func (m *MockDb) GetById(arg0 int) database.Student {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(database.Student)
	return ret0
}

// GetById indicates an expected call of GetById.
func (mr *MockDbMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockDb)(nil).GetById), arg0)
}

// IsPass mocks base method.
func (m *MockDb) IsPass(arg0 database.Student) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPass", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPass indicates an expected call of IsPass.
func (mr *MockDbMockRecorder) IsPass(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPass", reflect.TypeOf((*MockDb)(nil).IsPass), arg0)
}

// Set mocks base method.
func (m *MockDb) Set(arg0 database.Student) database.Class {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0)
	ret0, _ := ret[0].(database.Class)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockDbMockRecorder) Set(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDb)(nil).Set), arg0)
}
