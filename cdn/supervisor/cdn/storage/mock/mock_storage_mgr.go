// Code generated by MockGen. DO NOT EDIT.
// Source: d7y.io/dragonfly/v2/cdn/supervisor/cdn/storage (interfaces: Manager)

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	storedriver "d7y.io/dragonfly/v2/cdn/storedriver"
	supervisor "d7y.io/dragonfly/v2/cdn/supervisor"
	storage "d7y.io/dragonfly/v2/cdn/supervisor/cdn/storage"
	types "d7y.io/dragonfly/v2/cdn/types"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AppendPieceMetadata mocks base method.
func (m *MockManager) AppendPieceMetadata(arg0 string, arg1 *storage.PieceMetaRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendPieceMetadata", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendPieceMetadata indicates an expected call of AppendPieceMetadata.
func (mr *MockManagerMockRecorder) AppendPieceMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendPieceMetadata", reflect.TypeOf((*MockManager)(nil).AppendPieceMetadata), arg0, arg1)
}

// CreateUploadLink mocks base method.
func (m *MockManager) CreateUploadLink(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUploadLink", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUploadLink indicates an expected call of CreateUploadLink.
func (mr *MockManagerMockRecorder) CreateUploadLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUploadLink", reflect.TypeOf((*MockManager)(nil).CreateUploadLink), arg0)
}

// DeleteTask mocks base method.
func (m *MockManager) DeleteTask(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockManagerMockRecorder) DeleteTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockManager)(nil).DeleteTask), arg0)
}

// Initialize mocks base method.
func (m *MockManager) Initialize(arg0 supervisor.SeedTaskManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Initialize", arg0)
}

// Initialize indicates an expected call of Initialize.
func (mr *MockManagerMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockManager)(nil).Initialize), arg0)
}

// ReadDownloadFile mocks base method.
func (m *MockManager) ReadDownloadFile(arg0 string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDownloadFile", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDownloadFile indicates an expected call of ReadDownloadFile.
func (mr *MockManagerMockRecorder) ReadDownloadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDownloadFile", reflect.TypeOf((*MockManager)(nil).ReadDownloadFile), arg0)
}

// ReadFileMetadata mocks base method.
func (m *MockManager) ReadFileMetadata(arg0 string) (*storage.FileMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFileMetadata", arg0)
	ret0, _ := ret[0].(*storage.FileMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFileMetadata indicates an expected call of ReadFileMetadata.
func (mr *MockManagerMockRecorder) ReadFileMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFileMetadata", reflect.TypeOf((*MockManager)(nil).ReadFileMetadata), arg0)
}

// ReadPieceMetaRecords mocks base method.
func (m *MockManager) ReadPieceMetaRecords(arg0 string) ([]*storage.PieceMetaRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPieceMetaRecords", arg0)
	ret0, _ := ret[0].([]*storage.PieceMetaRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPieceMetaRecords indicates an expected call of ReadPieceMetaRecords.
func (mr *MockManagerMockRecorder) ReadPieceMetaRecords(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPieceMetaRecords", reflect.TypeOf((*MockManager)(nil).ReadPieceMetaRecords), arg0)
}

// ResetRepo mocks base method.
func (m *MockManager) ResetRepo(arg0 *types.SeedTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetRepo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetRepo indicates an expected call of ResetRepo.
func (mr *MockManagerMockRecorder) ResetRepo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetRepo", reflect.TypeOf((*MockManager)(nil).ResetRepo), arg0)
}

// StatDownloadFile mocks base method.
func (m *MockManager) StatDownloadFile(arg0 string) (*storedriver.StorageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatDownloadFile", arg0)
	ret0, _ := ret[0].(*storedriver.StorageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatDownloadFile indicates an expected call of StatDownloadFile.
func (mr *MockManagerMockRecorder) StatDownloadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatDownloadFile", reflect.TypeOf((*MockManager)(nil).StatDownloadFile), arg0)
}

// TryFreeSpace mocks base method.
func (m *MockManager) TryFreeSpace(arg0 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryFreeSpace", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TryFreeSpace indicates an expected call of TryFreeSpace.
func (mr *MockManagerMockRecorder) TryFreeSpace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryFreeSpace", reflect.TypeOf((*MockManager)(nil).TryFreeSpace), arg0)
}

// WriteDownloadFile mocks base method.
func (m *MockManager) WriteDownloadFile(arg0 string, arg1, arg2 int64, arg3 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteDownloadFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteDownloadFile indicates an expected call of WriteDownloadFile.
func (mr *MockManagerMockRecorder) WriteDownloadFile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteDownloadFile", reflect.TypeOf((*MockManager)(nil).WriteDownloadFile), arg0, arg1, arg2, arg3)
}

// WriteFileMetadata mocks base method.
func (m *MockManager) WriteFileMetadata(arg0 string, arg1 *storage.FileMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFileMetadata", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFileMetadata indicates an expected call of WriteFileMetadata.
func (mr *MockManagerMockRecorder) WriteFileMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFileMetadata", reflect.TypeOf((*MockManager)(nil).WriteFileMetadata), arg0, arg1)
}

// WritePieceMetaRecords mocks base method.
func (m *MockManager) WritePieceMetaRecords(arg0 string, arg1 []*storage.PieceMetaRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePieceMetaRecords", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePieceMetaRecords indicates an expected call of WritePieceMetaRecords.
func (mr *MockManagerMockRecorder) WritePieceMetaRecords(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePieceMetaRecords", reflect.TypeOf((*MockManager)(nil).WritePieceMetaRecords), arg0, arg1)
}
