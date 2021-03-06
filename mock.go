// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package main is a generated GoMock package.
package main

import (
	io "io"
	os "os"
	reflect "reflect"

	toml "github.com/BurntSushi/toml"
	gomock "github.com/golang/mock/gomock"
)

// MockhttpClient is a mock of httpClient interface.
type MockhttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockhttpClientMockRecorder
}

// MockhttpClientMockRecorder is the mock recorder for MockhttpClient.
type MockhttpClientMockRecorder struct {
	mock *MockhttpClient
}

// NewMockhttpClient creates a new mock instance.
func NewMockhttpClient(ctrl *gomock.Controller) *MockhttpClient {
	mock := &MockhttpClient{ctrl: ctrl}
	mock.recorder = &MockhttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockhttpClient) EXPECT() *MockhttpClientMockRecorder {
	return m.recorder
}

// post mocks base method.
func (m *MockhttpClient) post(path string, value io.Reader, header map[string]string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "post", path, value, header)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// post indicates an expected call of post.
func (mr *MockhttpClientMockRecorder) post(path, value, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "post", reflect.TypeOf((*MockhttpClient)(nil).post), path, value, header)
}

// MockioHandler is a mock of ioHandler interface.
type MockioHandler struct {
	ctrl     *gomock.Controller
	recorder *MockioHandlerMockRecorder
}

// MockioHandlerMockRecorder is the mock recorder for MockioHandler.
type MockioHandlerMockRecorder struct {
	mock *MockioHandler
}

// NewMockioHandler creates a new mock instance.
func NewMockioHandler(ctrl *gomock.Controller) *MockioHandler {
	mock := &MockioHandler{ctrl: ctrl}
	mock.recorder = &MockioHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockioHandler) EXPECT() *MockioHandlerMockRecorder {
	return m.recorder
}

// DecodeToml mocks base method.
func (m *MockioHandler) DecodeToml(data string, v interface{}) (toml.MetaData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeToml", data, v)
	ret0, _ := ret[0].(toml.MetaData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeToml indicates an expected call of DecodeToml.
func (mr *MockioHandlerMockRecorder) DecodeToml(data, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeToml", reflect.TypeOf((*MockioHandler)(nil).DecodeToml), data, v)
}

// GetHomeDirPath mocks base method.
func (m *MockioHandler) GetHomeDirPath() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHomeDirPath")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHomeDirPath indicates an expected call of GetHomeDirPath.
func (mr *MockioHandlerMockRecorder) GetHomeDirPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHomeDirPath", reflect.TypeOf((*MockioHandler)(nil).GetHomeDirPath))
}

// MakeDir mocks base method.
func (m *MockioHandler) MakeDir(dirPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDir", dirPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir.
func (mr *MockioHandlerMockRecorder) MakeDir(dirPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockioHandler)(nil).MakeDir), dirPath)
}

// NotExists mocks base method.
func (m *MockioHandler) NotExists(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotExists", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// NotExists indicates an expected call of NotExists.
func (mr *MockioHandlerMockRecorder) NotExists(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotExists", reflect.TypeOf((*MockioHandler)(nil).NotExists), path)
}

// OpenFile mocks base method.
func (m *MockioHandler) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", name, flag, perm)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile.
func (mr *MockioHandlerMockRecorder) OpenFile(name, flag, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockioHandler)(nil).OpenFile), name, flag, perm)
}

// ReadFile mocks base method.
func (m *MockioHandler) ReadFile(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockioHandlerMockRecorder) ReadFile(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockioHandler)(nil).ReadFile), path)
}

// RemoveFile mocks base method.
func (m *MockioHandler) RemoveFile(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFile", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFile indicates an expected call of RemoveFile.
func (mr *MockioHandlerMockRecorder) RemoveFile(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFile", reflect.TypeOf((*MockioHandler)(nil).RemoveFile), path)
}

// Write mocks base method.
func (m *MockioHandler) Write(f *os.File, b []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", f, b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockioHandlerMockRecorder) Write(f, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockioHandler)(nil).Write), f, b)
}
