// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// ImageController is an autogenerated mock type for the ImageController type
type ImageController struct {
	mock.Mock
}

// DownloadImage provides a mock function with given fields: ctx
func (_m *ImageController) DownloadImage(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadImage provides a mock function with given fields: ctx
func (_m *ImageController) UploadImage(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewImageController interface {
	mock.TestingT
	Cleanup(func())
}

// NewImageController creates a new instance of ImageController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImageController(t mockConstructorTestingTNewImageController) *ImageController {
	mock := &ImageController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}