// Code generated by MockGen. DO NOT EDIT.
// Source: notifications.go

// Package machine is a generated GoMock package.
package machine

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNotificationService is a mock of NotificationService interface
type MockNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceMockRecorder
}

// MockNotificationServiceMockRecorder is the mock recorder for MockNotificationService
type MockNotificationServiceMockRecorder struct {
	mock *MockNotificationService
}

// NewMockNotificationService creates a new mock instance
func NewMockNotificationService(ctrl *gomock.Controller) *MockNotificationService {
	mock := &MockNotificationService{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotificationService) EXPECT() *MockNotificationServiceMockRecorder {
	return m.recorder
}

// NotifyPostTransition mocks base method
func (m *MockNotificationService) NotifyPostTransition(t *TransitionNotification) error {
	ret := m.ctrl.Call(m, "NotifyPostTransition", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyPostTransition indicates an expected call of NotifyPostTransition
func (mr *MockNotificationServiceMockRecorder) NotifyPostTransition(t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyPostTransition", reflect.TypeOf((*MockNotificationService)(nil).NotifyPostTransition), t)
}

// NotifyWatcherState mocks base method
func (m *MockNotificationService) NotifyWatcherState(machine, watcher string, state map[string]interface{}) error {
	ret := m.ctrl.Call(m, "NotifyWatcherState", machine, watcher, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyWatcherState indicates an expected call of NotifyWatcherState
func (mr *MockNotificationServiceMockRecorder) NotifyWatcherState(machine, watcher, state interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWatcherState", reflect.TypeOf((*MockNotificationService)(nil).NotifyWatcherState), machine, watcher, state)
}

// Debugf mocks base method
func (m *MockNotificationService) Debugf(machine, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockNotificationServiceMockRecorder) Debugf(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockNotificationService)(nil).Debugf), varargs...)
}

// Infof mocks base method
func (m *MockNotificationService) Infof(machine, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockNotificationServiceMockRecorder) Infof(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockNotificationService)(nil).Infof), varargs...)
}

// Warnf mocks base method
func (m *MockNotificationService) Warnf(machine, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf
func (mr *MockNotificationServiceMockRecorder) Warnf(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockNotificationService)(nil).Warnf), varargs...)
}

// Errorf mocks base method
func (m *MockNotificationService) Errorf(machine, watcher, format string, args ...interface{}) {
	varargs := []interface{}{machine, watcher, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockNotificationServiceMockRecorder) Errorf(machine, watcher, format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{machine, watcher, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockNotificationService)(nil).Errorf), varargs...)
}