// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// BookController is an autogenerated mock type for the BookController type
type BookController struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: _a0
func (_m *BookController) CreateBook(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: _a0
func (_m *BookController) DeleteBook(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBook provides a mock function with given fields: _a0
func (_m *BookController) GetBook(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListBooks provides a mock function with given fields: _a0
func (_m *BookController) ListBooks(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBook provides a mock function with given fields: _a0
func (_m *BookController) UpdateBook(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBookController interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookController creates a new instance of BookController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookController(t mockConstructorTestingTNewBookController) *BookController {
	mock := &BookController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
