// Code generated by MockGen. DO NOT EDIT.
// Source: internal/interfaces/app.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/sveboo/exchangerate/internal/models"
)

// MockApp is a mock of App interface.
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *MockAppMockRecorder
}

// MockAppMockRecorder is the mock recorder for MockApp.
type MockAppMockRecorder struct {
	mock *MockApp
}

// NewMockApp creates a new mock instance.
func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &MockAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApp) EXPECT() *MockAppMockRecorder {
	return m.recorder
}

// GetCurrencyRate mocks base method.
func (m *MockApp) GetCurrencyRate(date, currency string) (map[string]float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencyRate", date, currency)
	ret0, _ := ret[0].(map[string]float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencyRate indicates an expected call of GetCurrencyRate.
func (mr *MockAppMockRecorder) GetCurrencyRate(date, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencyRate", reflect.TypeOf((*MockApp)(nil).GetCurrencyRate), date, currency)
}

// GetInfo mocks base method.
func (m *MockApp) GetInfo() models.InfoApp {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(models.InfoApp)
	return ret0
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockAppMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockApp)(nil).GetInfo))
}

// IsCurrencyValid mocks base method.
func (m *MockApp) IsCurrencyValid(currency string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCurrencyValid", currency)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCurrencyValid indicates an expected call of IsCurrencyValid.
func (mr *MockAppMockRecorder) IsCurrencyValid(currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCurrencyValid", reflect.TypeOf((*MockApp)(nil).IsCurrencyValid), currency)
}
