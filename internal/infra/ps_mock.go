// Code generated by MockGen. DO NOT EDIT.
// Source: ps.go
//
// Generated by this command:
//
//	mockgen -source=ps.go -destination=ps_mock.go -package=infra
//

// Package infra is a generated GoMock package.
package infra

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPsInterface is a mock of PsInterface interface.
type MockPsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPsInterfaceMockRecorder
	isgomock struct{}
}

// MockPsInterfaceMockRecorder is the mock recorder for MockPsInterface.
type MockPsInterfaceMockRecorder struct {
	mock *MockPsInterface
}

// NewMockPsInterface creates a new mock instance.
func NewMockPsInterface(ctrl *gomock.Controller) *MockPsInterface {
	mock := &MockPsInterface{ctrl: ctrl}
	mock.recorder = &MockPsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPsInterface) EXPECT() *MockPsInterfaceMockRecorder {
	return m.recorder
}

// Exit mocks base method.
func (m *MockPsInterface) Exit(code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Exit", code)
}

// Exit indicates an expected call of Exit.
func (mr *MockPsInterfaceMockRecorder) Exit(code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exit", reflect.TypeOf((*MockPsInterface)(nil).Exit), code)
}

// Print mocks base method.
func (m *MockPsInterface) Print(text string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Print", text)
}

// Print indicates an expected call of Print.
func (mr *MockPsInterfaceMockRecorder) Print(text any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockPsInterface)(nil).Print), text)
}

// PrintErr mocks base method.
func (m *MockPsInterface) PrintErr(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrintErr", err)
}

// PrintErr indicates an expected call of PrintErr.
func (mr *MockPsInterfaceMockRecorder) PrintErr(err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintErr", reflect.TypeOf((*MockPsInterface)(nil).PrintErr), err)
}

// Printf mocks base method.
func (m *MockPsInterface) Printf(format string, a ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf.
func (mr *MockPsInterfaceMockRecorder) Printf(format any, a ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, a...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockPsInterface)(nil).Printf), varargs...)
}
