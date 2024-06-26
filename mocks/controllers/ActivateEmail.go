// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// ActivateEmail is an autogenerated mock type for the ActivateEmail type
type ActivateEmail struct {
	mock.Mock
}

// ActivateEmail provides a mock function with given fields: c
func (_m *ActivateEmail) ActivateEmail(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ActivateEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewActivateEmail creates a new instance of ActivateEmail. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewActivateEmail(t interface {
	mock.TestingT
	Cleanup(func())
}) *ActivateEmail {
	mock := &ActivateEmail{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
