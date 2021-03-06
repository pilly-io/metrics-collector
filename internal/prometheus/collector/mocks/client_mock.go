// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pilly-io/metrics-collector/internal/prometheus/client (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/pilly-io/metrics-collector/internal/prometheus/client"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetPodsCPURequests mocks base method
func (m *MockClient) GetPodsCPURequests(arg0 string) (client.MetricsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodsCPURequests", arg0)
	ret0, _ := ret[0].(client.MetricsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodsCPURequests indicates an expected call of GetPodsCPURequests
func (mr *MockClientMockRecorder) GetPodsCPURequests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodsCPURequests", reflect.TypeOf((*MockClient)(nil).GetPodsCPURequests), arg0)
}

// GetPodsCPUUsage mocks base method
func (m *MockClient) GetPodsCPUUsage(arg0 string) (client.MetricsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodsCPUUsage", arg0)
	ret0, _ := ret[0].(client.MetricsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodsCPUUsage indicates an expected call of GetPodsCPUUsage
func (mr *MockClientMockRecorder) GetPodsCPUUsage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodsCPUUsage", reflect.TypeOf((*MockClient)(nil).GetPodsCPUUsage), arg0)
}

// GetPodsMemoryRequests mocks base method
func (m *MockClient) GetPodsMemoryRequests(arg0 string) (client.MetricsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodsMemoryRequests", arg0)
	ret0, _ := ret[0].(client.MetricsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodsMemoryRequests indicates an expected call of GetPodsMemoryRequests
func (mr *MockClientMockRecorder) GetPodsMemoryRequests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodsMemoryRequests", reflect.TypeOf((*MockClient)(nil).GetPodsMemoryRequests), arg0)
}

// GetPodsMemoryUsage mocks base method
func (m *MockClient) GetPodsMemoryUsage(arg0 string) (client.MetricsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodsMemoryUsage", arg0)
	ret0, _ := ret[0].(client.MetricsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodsMemoryUsage indicates an expected call of GetPodsMemoryUsage
func (mr *MockClientMockRecorder) GetPodsMemoryUsage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodsMemoryUsage", reflect.TypeOf((*MockClient)(nil).GetPodsMemoryUsage), arg0)
}
