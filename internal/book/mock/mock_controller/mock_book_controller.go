// Code generated by MockGen. DO NOT EDIT.
// Source: book_controller.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockBookController is a mock of BookController interface
type MockBookController struct {
	ctrl     *gomock.Controller
	recorder *MockBookControllerMockRecorder
}

// MockBookControllerMockRecorder is the mock recorder for MockBookController
type MockBookControllerMockRecorder struct {
	mock *MockBookController
}

// NewMockBookController creates a new mock instance
func NewMockBookController(ctrl *gomock.Controller) *MockBookController {
	mock := &MockBookController{ctrl: ctrl}
	mock.recorder = &MockBookControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookController) EXPECT() *MockBookControllerMockRecorder {
	return m.recorder
}

// GetBook mocks base method
func (m *MockBookController) GetBook(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBook", w, r)
}

// GetBook indicates an expected call of GetBook
func (mr *MockBookControllerMockRecorder) GetBook(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBook", reflect.TypeOf((*MockBookController)(nil).GetBook), w, r)
}
