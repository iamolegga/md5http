// Code generated by MockGen. DO NOT EDIT.
// Source: io (interfaces: StringWriter)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStringWriter is a mock of StringWriter interface.
type MockStringWriter struct {
	ctrl     *gomock.Controller
	recorder *MockStringWriterMockRecorder
}

// MockStringWriterMockRecorder is the mock recorder for MockStringWriter.
type MockStringWriterMockRecorder struct {
	mock *MockStringWriter
}

// NewMockStringWriter creates a new mock instance.
func NewMockStringWriter(ctrl *gomock.Controller) *MockStringWriter {
	mock := &MockStringWriter{ctrl: ctrl}
	mock.recorder = &MockStringWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStringWriter) EXPECT() *MockStringWriterMockRecorder {
	return m.recorder
}

// WriteString mocks base method.
func (m *MockStringWriter) WriteString(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString.
func (mr *MockStringWriterMockRecorder) WriteString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockStringWriter)(nil).WriteString), arg0)
}
