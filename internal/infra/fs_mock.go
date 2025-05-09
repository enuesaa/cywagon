// Code generated by MockGen. DO NOT EDIT.
// Source: fs.go
//
// Generated by this command:
//
//	mockgen -source=fs.go -destination=fs_mock.go -package=infra
//

// Package infra is a generated GoMock package.
package infra

import (
	fs "io/fs"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFsInterface is a mock of FsInterface interface.
type MockFsInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFsInterfaceMockRecorder
	isgomock struct{}
}

// MockFsInterfaceMockRecorder is the mock recorder for MockFsInterface.
type MockFsInterfaceMockRecorder struct {
	mock *MockFsInterface
}

// NewMockFsInterface creates a new mock instance.
func NewMockFsInterface(ctrl *gomock.Controller) *MockFsInterface {
	mock := &MockFsInterface{ctrl: ctrl}
	mock.recorder = &MockFsInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFsInterface) EXPECT() *MockFsInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFsInterface) Create(path string, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", path, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockFsInterfaceMockRecorder) Create(path, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFsInterface)(nil).Create), path, body)
}

// DirFS mocks base method.
func (m *MockFsInterface) DirFS(path string) (fs.FS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DirFS", path)
	ret0, _ := ret[0].(fs.FS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DirFS indicates an expected call of DirFS.
func (mr *MockFsInterfaceMockRecorder) DirFS(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DirFS", reflect.TypeOf((*MockFsInterface)(nil).DirFS), path)
}

// IsExist mocks base method.
func (m *MockFsInterface) IsExist(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExist indicates an expected call of IsExist.
func (mr *MockFsInterfaceMockRecorder) IsExist(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockFsInterface)(nil).IsExist), path)
}

// IsFile mocks base method.
func (m *MockFsInterface) IsFile(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFile", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFile indicates an expected call of IsFile.
func (mr *MockFsInterfaceMockRecorder) IsFile(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFile", reflect.TypeOf((*MockFsInterface)(nil).IsFile), path)
}

// ListFiles mocks base method.
func (m *MockFsInterface) ListFiles(path string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", path)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockFsInterfaceMockRecorder) ListFiles(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockFsInterface)(nil).ListFiles), path)
}

// Read mocks base method.
func (m *MockFsInterface) Read(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockFsInterfaceMockRecorder) Read(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFsInterface)(nil).Read), path)
}
