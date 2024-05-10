// Code generated by MockGen. DO NOT EDIT.
// Source: internal/interfaces/parser.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/sveboo/exchangerate/internal/models"
)

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser.
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockParser) Parse(data io.ReadCloser) (map[string]float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", data)
	ret0, _ := ret[0].(map[string]float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockParserMockRecorder) Parse(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), data)
}

// ParseCurrencyISO mocks base method.
func (m *MockParser) ParseCurrencyISO(data io.ReadCloser) ([]models.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseCurrencyISO", data)
	ret0, _ := ret[0].([]models.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseCurrencyISO indicates an expected call of ParseCurrencyISO.
func (mr *MockParserMockRecorder) ParseCurrencyISO(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseCurrencyISO", reflect.TypeOf((*MockParser)(nil).ParseCurrencyISO), data)
}