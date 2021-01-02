// Code generated by MockGen. DO NOT EDIT.
// Source: book_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/zhendong233/Books/internal/book/model"
	reflect "reflect"
)

// MockBookService is a mock of BookService interface
type MockBookService struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceMockRecorder
}

// MockBookServiceMockRecorder is the mock recorder for MockBookService
type MockBookServiceMockRecorder struct {
	mock *MockBookService
}

// NewMockBookService creates a new mock instance
func NewMockBookService(ctrl *gomock.Controller) *MockBookService {
	mock := &MockBookService{ctrl: ctrl}
	mock.recorder = &MockBookServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookService) EXPECT() *MockBookServiceMockRecorder {
	return m.recorder
}

// FindByID mocks base method
func (m *MockBookService) FindByID(ctx context.Context, bookID string) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, bookID)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockBookServiceMockRecorder) FindByID(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockBookService)(nil).FindByID), ctx, bookID)
}

// CreateBook mocks base method
func (m *MockBookService) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", ctx, book)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook
func (mr *MockBookServiceMockRecorder) CreateBook(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockBookService)(nil).CreateBook), ctx, book)
}

// UpdateBook mocks base method
func (m *MockBookService) UpdateBook(ctx context.Context, bookID string, book *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", ctx, bookID, book)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook
func (mr *MockBookServiceMockRecorder) UpdateBook(ctx, bookID, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockBookService)(nil).UpdateBook), ctx, bookID, book)
}
