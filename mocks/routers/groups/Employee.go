// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// Employee is an autogenerated mock type for the Employee type
type Employee struct {
	mock.Mock
}

// Resource provides a mock function with given fields: g
func (_m *Employee) Resource(g *echo.Group) {
	_m.Called(g)
}

// NewEmployee creates a new instance of Employee. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmployee(t interface {
	mock.TestingT
	Cleanup(func())
}) *Employee {
	mock := &Employee{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
