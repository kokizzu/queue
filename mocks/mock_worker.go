// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/golang-queue/queue/core (interfaces: Worker)
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=mock_worker.go github.com/golang-queue/queue/core Worker
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	core "github.com/golang-queue/queue/core"
	gomock "go.uber.org/mock/gomock"
)

// MockWorker is a mock of Worker interface.
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker.
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance.
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}

// Queue mocks base method.
func (m *MockWorker) Queue(arg0 core.TaskMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Queue", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Queue indicates an expected call of Queue.
func (mr *MockWorkerMockRecorder) Queue(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Queue", reflect.TypeOf((*MockWorker)(nil).Queue), arg0)
}

// Request mocks base method.
func (m *MockWorker) Request() (core.TaskMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(core.TaskMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockWorkerMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockWorker)(nil).Request))
}

// Run mocks base method.
func (m *MockWorker) Run(arg0 context.Context, arg1 core.TaskMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockWorkerMockRecorder) Run(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWorker)(nil).Run), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockWorker) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockWorkerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockWorker)(nil).Shutdown))
}
