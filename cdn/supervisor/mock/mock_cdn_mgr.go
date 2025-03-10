// Code generated by MockGen. DO NOT EDIT.
// Source: d7y.io/dragonfly/v2/cdn/supervisor (interfaces: CDNMgr)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	types "d7y.io/dragonfly/v2/cdn/types"
	gomock "github.com/golang/mock/gomock"
)

// MockCDNMgr is a mock of CDNMgr interface.
type MockCDNMgr struct {
	ctrl     *gomock.Controller
	recorder *MockCDNMgrMockRecorder
}

// MockCDNMgrMockRecorder is the mock recorder for MockCDNMgr.
type MockCDNMgrMockRecorder struct {
	mock *MockCDNMgr
}

// NewMockCDNMgr creates a new mock instance.
func NewMockCDNMgr(ctrl *gomock.Controller) *MockCDNMgr {
	mock := &MockCDNMgr{ctrl: ctrl}
	mock.recorder = &MockCDNMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCDNMgr) EXPECT() *MockCDNMgrMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCDNMgr) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCDNMgrMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCDNMgr)(nil).Delete), arg0)
}

// TriggerCDN mocks base method.
func (m *MockCDNMgr) TriggerCDN(arg0 context.Context, arg1 *types.SeedTask) (*types.SeedTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerCDN", arg0, arg1)
	ret0, _ := ret[0].(*types.SeedTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerCDN indicates an expected call of TriggerCDN.
func (mr *MockCDNMgrMockRecorder) TriggerCDN(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerCDN", reflect.TypeOf((*MockCDNMgr)(nil).TriggerCDN), arg0, arg1)
}

// TryFreeSpace mocks base method.
func (m *MockCDNMgr) TryFreeSpace(arg0 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryFreeSpace", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TryFreeSpace indicates an expected call of TryFreeSpace.
func (mr *MockCDNMgrMockRecorder) TryFreeSpace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryFreeSpace", reflect.TypeOf((*MockCDNMgr)(nil).TryFreeSpace), arg0)
}
