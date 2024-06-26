// Code generated by MockGen. DO NOT EDIT.
// Source: internal/interfaces/currency_servicer.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCurrencyServicer is a mock of CurrencyServicer interface.
type MockCurrencyServicer struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyServicerMockRecorder
}

// MockCurrencyServicerMockRecorder is the mock recorder for MockCurrencyServicer.
type MockCurrencyServicerMockRecorder struct {
	mock *MockCurrencyServicer
}

// NewMockCurrencyServicer creates a new mock instance.
func NewMockCurrencyServicer(ctrl *gomock.Controller) *MockCurrencyServicer {
	mock := &MockCurrencyServicer{ctrl: ctrl}
	mock.recorder = &MockCurrencyServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyServicer) EXPECT() *MockCurrencyServicerMockRecorder {
	return m.recorder
}

// GetCurrencyInfo mocks base method.
func (m *MockCurrencyServicer) GetCurrencyInfo() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencyInfo")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencyInfo indicates an expected call of GetCurrencyInfo.
func (mr *MockCurrencyServicerMockRecorder) GetCurrencyInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencyInfo", reflect.TypeOf((*MockCurrencyServicer)(nil).GetCurrencyInfo))
}

// GetExchangeRate mocks base method.
func (m *MockCurrencyServicer) GetExchangeRate(date string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", date)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate.
func (mr *MockCurrencyServicerMockRecorder) GetExchangeRate(date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockCurrencyServicer)(nil).GetExchangeRate), date)
}
