// Code generated by MockGen. DO NOT EDIT.
// Source: ./estimator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entities "mercadolibre.com/di/practice/entities"
	reflect "reflect"
)

// MockweatherFetcher is a mock of weatherFetcher interface
type MockweatherFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockweatherFetcherMockRecorder
}

// MockweatherFetcherMockRecorder is the mock recorder for MockweatherFetcher
type MockweatherFetcherMockRecorder struct {
	mock *MockweatherFetcher
}

// NewMockweatherFetcher creates a new mock instance
func NewMockweatherFetcher(ctrl *gomock.Controller) *MockweatherFetcher {
	mock := &MockweatherFetcher{ctrl: ctrl}
	mock.recorder = &MockweatherFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockweatherFetcher) EXPECT() *MockweatherFetcherMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockweatherFetcher) Get(country, state, city string, forecastDays uint) (*entities.Forecast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", country, state, city, forecastDays)
	ret0, _ := ret[0].(*entities.Forecast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockweatherFetcherMockRecorder) Get(country, state, city, forecastDays interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockweatherFetcher)(nil).Get), country, state, city, forecastDays)
}