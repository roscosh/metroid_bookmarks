// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/sql/photos/photos.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/sql/photos/photos.go -destination=./internal/repository/sql/photos/mocks/photos.go
//

// Package mock_photos is a generated GoMock package.
package mock_photos

import (
	photos "metroid_bookmarks/internal/repository/sql/photos"
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
func (m *MockSQL) Create(createForm *photos.CreatePhoto) (*photos.PhotoPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", createForm)
	ret0, _ := ret[0].(*photos.PhotoPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSQLMockRecorder) Create(createForm any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSQL)(nil).Create), createForm)
}

// Delete mocks base method.
func (m *MockSQL) Delete(photoID, userID int) (*photos.PhotoPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", photoID, userID)
	ret0, _ := ret[0].(*photos.PhotoPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSQLMockRecorder) Delete(photoID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSQL)(nil).Delete), photoID, userID)
}
