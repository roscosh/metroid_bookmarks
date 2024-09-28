// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/pgpool/sql.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/pgpool/sql.go -destination=./pkg/pgpool/mocks/sql.go
//

// Package mock_pgpool is a generated GoMock package.
package mock_pgpool

import (
	context "context"
	reflect "reflect"

	pgx "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	gomock "go.uber.org/mock/gomock"
)

// MockSQL is a mock of SQL interface.
type MockSQL[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockSQLMockRecorder[T]
}

// MockSQLMockRecorder is the mock recorder for MockSQL.
type MockSQLMockRecorder[T any] struct {
	mock *MockSQL[T]
}

// NewMockSQL creates a new mock instance.
func NewMockSQL[T any](ctrl *gomock.Controller) *MockSQL[T] {
	mock := &MockSQL[T]{ctrl: ctrl}
	mock.recorder = &MockSQLMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQL[T]) EXPECT() *MockSQLMockRecorder[T] {
	return m.recorder
}

// CollectOneRow mocks base method.
func (m *MockSQL[T]) CollectOneRow(rows pgx.Rows) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectOneRow", rows)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectOneRow indicates an expected call of CollectOneRow.
func (mr *MockSQLMockRecorder[T]) CollectOneRow(rows any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectOneRow", reflect.TypeOf((*MockSQL[T])(nil).CollectOneRow), rows)
}

// CollectRows mocks base method.
func (m *MockSQL[T]) CollectRows(rows pgx.Rows) ([]T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectRows", rows)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectRows indicates an expected call of CollectRows.
func (mr *MockSQLMockRecorder[T]) CollectRows(rows any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectRows", reflect.TypeOf((*MockSQL[T])(nil).CollectRows), rows)
}

// Delete mocks base method.
func (m *MockSQL[T]) Delete(ctx context.Context, pk int) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, pk)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSQLMockRecorder[T]) Delete(ctx, pk any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSQL[T])(nil).Delete), ctx, pk)
}

// DeleteWhere mocks base method.
func (m *MockSQL[T]) DeleteWhere(ctx context.Context, whereStatement string, args ...any) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, whereStatement}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWhere", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWhere indicates an expected call of DeleteWhere.
func (mr *MockSQLMockRecorder[T]) DeleteWhere(ctx, whereStatement any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, whereStatement}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWhere", reflect.TypeOf((*MockSQL[T])(nil).DeleteWhere), varargs...)
}

// Exec mocks base method.
func (m *MockSQL[T]) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockSQLMockRecorder[T]) Exec(ctx, query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockSQL[T])(nil).Exec), varargs...)
}

// Insert mocks base method.
func (m *MockSQL[T]) Insert(ctx context.Context, createStruct any) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, createStruct)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockSQLMockRecorder[T]) Insert(ctx, createStruct any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSQL[T])(nil).Insert), ctx, createStruct)
}

// Query mocks base method.
func (m *MockSQL[T]) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockSQLMockRecorder[T]) Query(ctx, query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSQL[T])(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockSQL[T]) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []any{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockSQLMockRecorder[T]) QueryRow(ctx, query any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockSQL[T])(nil).QueryRow), varargs...)
}

// Select mocks base method.
func (m *MockSQL[T]) Select(ctx context.Context, pk int) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx, pk)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockSQLMockRecorder[T]) Select(ctx, pk any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockSQL[T])(nil).Select), ctx, pk)
}

// SelectMany mocks base method.
func (m *MockSQL[T]) SelectMany(ctx context.Context) ([]T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMany", ctx)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMany indicates an expected call of SelectMany.
func (mr *MockSQLMockRecorder[T]) SelectMany(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMany", reflect.TypeOf((*MockSQL[T])(nil).SelectMany), ctx)
}

// SelectManyWhere mocks base method.
func (m *MockSQL[T]) SelectManyWhere(ctx context.Context, whereStatement string, args ...any) ([]T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, whereStatement}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectManyWhere", varargs...)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectManyWhere indicates an expected call of SelectManyWhere.
func (mr *MockSQLMockRecorder[T]) SelectManyWhere(ctx, whereStatement any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, whereStatement}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectManyWhere", reflect.TypeOf((*MockSQL[T])(nil).SelectManyWhere), varargs...)
}

// SelectWhere mocks base method.
func (m *MockSQL[T]) SelectWhere(ctx context.Context, whereStatement string, args ...any) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, whereStatement}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectWhere", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectWhere indicates an expected call of SelectWhere.
func (mr *MockSQLMockRecorder[T]) SelectWhere(ctx, whereStatement any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, whereStatement}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectWhere", reflect.TypeOf((*MockSQL[T])(nil).SelectWhere), varargs...)
}

// Total mocks base method.
func (m *MockSQL[T]) Total(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Total", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Total indicates an expected call of Total.
func (mr *MockSQLMockRecorder[T]) Total(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Total", reflect.TypeOf((*MockSQL[T])(nil).Total), ctx)
}

// Update mocks base method.
func (m *MockSQL[T]) Update(ctx context.Context, pk int, editStruct any) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, pk, editStruct)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSQLMockRecorder[T]) Update(ctx, pk, editStruct any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSQL[T])(nil).Update), ctx, pk, editStruct)
}

// UpdateWhere mocks base method.
func (m *MockSQL[T]) UpdateWhere(ctx context.Context, editStruct any, where string, args ...any) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, editStruct, where}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateWhere", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWhere indicates an expected call of UpdateWhere.
func (mr *MockSQLMockRecorder[T]) UpdateWhere(ctx, editStruct, where any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, editStruct, where}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWhere", reflect.TypeOf((*MockSQL[T])(nil).UpdateWhere), varargs...)
}
