// Code generated by MockGen. DO NOT EDIT.
// Source: products_source.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	"cache-ahead-task/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProductsSource is a mock of ProductsSource interface.
type MockProductsSource struct {
	ctrl     *gomock.Controller
	recorder *MockProductsSourceMockRecorder
}

// MockProductsSourceMockRecorder is the mock recorder for MockProductsSource.
type MockProductsSourceMockRecorder struct {
	mock *MockProductsSource
}

// NewMockProductsSource creates a new mock instance.
func NewMockProductsSource(ctrl *gomock.Controller) *MockProductsSource {
	mock := &MockProductsSource{ctrl: ctrl}
	mock.recorder = &MockProductsSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsSource) EXPECT() *MockProductsSourceMockRecorder {
	return m.recorder
}

// GetProducts mocks base method.
func (m *MockProductsSource) GetProducts() []model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts")
	ret0, _ := ret[0].([]model.Product)
	return ret0
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductsSourceMockRecorder) GetProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductsSource)(nil).GetProducts))
}