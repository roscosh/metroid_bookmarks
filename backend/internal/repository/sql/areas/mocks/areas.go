// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/sql/areas/areas.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/sql/areas/areas.go -destination=./internal/repository/sql/areas/mocks/areas.go
//

// Package mock_areas is a generated GoMock package.
package mock_areas

import (
	areas "metroid_bookmarks/internal/repository/sql/areas"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSQL is a mock of SQL interface.
type MockSQL struct {
	ctrl     *gomock.Controller
	recorder *MockSQLMockRecorder
}

// MockSQLMockRecorder is the mock recorder for MockSQL.
type MockSQLMockRecorder struct {
	mock *MockSQL
}

// NewMockSQL creates a new mock instance.
func NewMockSQL(ctrl *gomock.Controller) *MockSQL {
	mock := &MockSQL{ctrl: ctrl}
	mock.recorder = &MockSQLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQL) EXPECT() *MockSQLMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSQL) Create(createForm *areas.CreateArea) (*areas.Area, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", createForm)
	ret0, _ := ret[0].(*areas.Area)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSQLMockRecorder) Create(createForm any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSQL)(nil).Create), createForm)
}

// Delete mocks base method.
func (m *MockSQL) Delete(id int) (*areas.Area, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(*areas.Area)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSQLMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSQL)(nil).Delete), id)
}

// Edit mocks base method.
func (m *MockSQL) Edit(id int, editForm *areas.EditArea) (*areas.Area, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", id, editForm)
	ret0, _ := ret[0].(*areas.Area)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Edit indicates an expected call of Edit.
func (mr *MockSQLMockRecorder) Edit(id, editForm any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockSQL)(nil).Edit), id, editForm)
}

// GetAll mocks base method.
func (m *MockSQL) GetAll() ([]areas.Area, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]areas.Area)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockSQLMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSQL)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockSQL) GetByID(id int) (*areas.Area, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*areas.Area)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockSQLMockRecorder) GetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockSQL)(nil).GetByID), id)
}

// Total mocks base method.
func (m *MockSQL) Total() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Total")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Total indicates an expected call of Total.
func (mr *MockSQLMockRecorder) Total() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Total", reflect.TypeOf((*MockSQL)(nil).Total))
}
