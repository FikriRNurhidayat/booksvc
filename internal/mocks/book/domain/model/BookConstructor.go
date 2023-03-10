// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	book_model "github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// BookConstructor is an autogenerated mock type for the BookConstructor type
type BookConstructor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *BookConstructor) Execute(_a0 book_model.Book) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewBookConstructor interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookConstructor creates a new instance of BookConstructor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookConstructor(t mockConstructorTestingTNewBookConstructor) *BookConstructor {
	mock := &BookConstructor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
