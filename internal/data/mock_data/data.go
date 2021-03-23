// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/ddns-updater/internal/data (interfaces: Database)

// Package mock_data is a generated GoMock package.
package mock_data

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/qdm12/ddns-updater/internal/models"
	records "github.com/qdm12/ddns-updater/internal/records"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDatabase) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaseMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabase)(nil).Close))
}

// GetEvents mocks base method.
func (m *MockDatabase) GetEvents(arg0, arg1 string) ([]models.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", arg0, arg1)
	ret0, _ := ret[0].([]models.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockDatabaseMockRecorder) GetEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockDatabase)(nil).GetEvents), arg0, arg1)
}

// Select mocks base method.
func (m *MockDatabase) Select(arg0 int) (records.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", arg0)
	ret0, _ := ret[0].(records.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockDatabaseMockRecorder) Select(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockDatabase)(nil).Select), arg0)
}

// SelectAll mocks base method.
func (m *MockDatabase) SelectAll() []records.Record {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]records.Record)
	return ret0
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockDatabaseMockRecorder) SelectAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockDatabase)(nil).SelectAll))
}

// Update mocks base method.
func (m *MockDatabase) Update(arg0 int, arg1 records.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDatabaseMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDatabase)(nil).Update), arg0, arg1)
}
