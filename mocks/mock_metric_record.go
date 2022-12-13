// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/ts-bridge/storage (interfaces: MetricRecord)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockMetricRecord is a mock of MetricRecord interface.
type MockMetricRecord struct {
	ctrl     *gomock.Controller
	recorder *MockMetricRecordMockRecorder
}

// MockMetricRecordMockRecorder is the mock recorder for MockMetricRecord.
type MockMetricRecordMockRecorder struct {
	mock *MockMetricRecord
}

// NewMockMetricRecord creates a new mock instance.
func NewMockMetricRecord(ctrl *gomock.Controller) *MockMetricRecord {
	mock := &MockMetricRecord{ctrl: ctrl}
	mock.recorder = &MockMetricRecordMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricRecord) EXPECT() *MockMetricRecordMockRecorder {
	return m.recorder
}

// GetCounterStartTime mocks base method.
func (m *MockMetricRecord) GetCounterStartTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounterStartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCounterStartTime indicates an expected call of GetCounterStartTime.
func (mr *MockMetricRecordMockRecorder) GetCounterStartTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounterStartTime", reflect.TypeOf((*MockMetricRecord)(nil).GetCounterStartTime))
}

// GetLastUpdate mocks base method.
func (m *MockMetricRecord) GetLastUpdate() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastUpdate")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLastUpdate indicates an expected call of GetLastUpdate.
func (mr *MockMetricRecordMockRecorder) GetLastUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastUpdate", reflect.TypeOf((*MockMetricRecord)(nil).GetLastUpdate))
}

// SetCounterStartTime mocks base method.
func (m *MockMetricRecord) SetCounterStartTime(arg0 context.Context, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCounterStartTime", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCounterStartTime indicates an expected call of SetCounterStartTime.
func (mr *MockMetricRecordMockRecorder) SetCounterStartTime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCounterStartTime", reflect.TypeOf((*MockMetricRecord)(nil).SetCounterStartTime), arg0, arg1)
}

// UpdateError mocks base method.
func (m *MockMetricRecord) UpdateError(arg0 context.Context, arg1 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateError indicates an expected call of UpdateError.
func (mr *MockMetricRecordMockRecorder) UpdateError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateError", reflect.TypeOf((*MockMetricRecord)(nil).UpdateError), arg0, arg1)
}

// UpdateSuccess mocks base method.
func (m *MockMetricRecord) UpdateSuccess(arg0 context.Context, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSuccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSuccess indicates an expected call of UpdateSuccess.
func (mr *MockMetricRecordMockRecorder) UpdateSuccess(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSuccess", reflect.TypeOf((*MockMetricRecord)(nil).UpdateSuccess), arg0, arg1, arg2)
}
