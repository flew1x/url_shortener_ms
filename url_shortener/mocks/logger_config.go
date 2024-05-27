// Code generated by MockGen. DO NOT EDIT.
// Source: internal/config/logger_config.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockILoggerConfig is a mock of ILoggerConfig interface.
type MockILoggerConfig struct {
	ctrl     *gomock.Controller
	recorder *MockILoggerConfigMockRecorder
}

// MockILoggerConfigMockRecorder is the mock recorder for MockILoggerConfig.
type MockILoggerConfigMockRecorder struct {
	mock *MockILoggerConfig
}

// NewMockILoggerConfig creates a new mock instance.
func NewMockILoggerConfig(ctrl *gomock.Controller) *MockILoggerConfig {
	mock := &MockILoggerConfig{ctrl: ctrl}
	mock.recorder = &MockILoggerConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILoggerConfig) EXPECT() *MockILoggerConfigMockRecorder {
	return m.recorder
}

// GetLogLevel mocks base method.
func (m *MockILoggerConfig) GetLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLogLevel indicates an expected call of GetLogLevel.
func (mr *MockILoggerConfigMockRecorder) GetLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogLevel", reflect.TypeOf((*MockILoggerConfig)(nil).GetLogLevel))
}
