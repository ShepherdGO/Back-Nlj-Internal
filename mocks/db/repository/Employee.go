// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	dtos "github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	mock "github.com/stretchr/testify/mock"

	models "github.com/Shepherd-Go/Back-Nlj-Internal.git/db/models"

	uuid "github.com/google/uuid"
)

// Employee is an autogenerated mock type for the Employee type
type Employee struct {
	mock.Mock
}

// CreateEmployee provides a mock function with given fields: ctx, empl
func (_m *Employee) CreateEmployee(ctx context.Context, empl models.Employee) error {
	ret := _m.Called(ctx, empl)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Employee) error); ok {
		r0 = rf(ctx, empl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEmployee provides a mock function with given fields: ctx, id, idToken
func (_m *Employee) DeleteEmployee(ctx context.Context, id uuid.UUID, idToken string) error {
	ret := _m.Called(ctx, id, idToken)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, id, idToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchAllEmployees provides a mock function with given fields: ctx
func (_m *Employee) SearchAllEmployees(ctx context.Context) (dtos.Employees, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SearchAllEmployees")
	}

	var r0 dtos.Employees
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (dtos.Employees, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) dtos.Employees); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dtos.Employees)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchEmployeByEmailOrUsername provides a mock function with given fields: ctx, identifier
func (_m *Employee) SearchEmployeByEmailOrUsername(ctx context.Context, identifier string) (dtos.EmployeeResponse, error) {
	ret := _m.Called(ctx, identifier)

	if len(ret) == 0 {
		panic("no return value specified for SearchEmployeByEmailOrUsername")
	}

	var r0 dtos.EmployeeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (dtos.EmployeeResponse, error)); ok {
		return rf(ctx, identifier)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) dtos.EmployeeResponse); ok {
		r0 = rf(ctx, identifier)
	} else {
		r0 = ret.Get(0).(dtos.EmployeeResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchEmployeeByEmail provides a mock function with given fields: ctx, email
func (_m *Employee) SearchEmployeeByEmail(ctx context.Context, email string) (dtos.EmployeeResponse, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for SearchEmployeeByEmail")
	}

	var r0 dtos.EmployeeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (dtos.EmployeeResponse, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) dtos.EmployeeResponse); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(dtos.EmployeeResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchEmployeeByEmailAndNotID provides a mock function with given fields: ctx, id, email
func (_m *Employee) SearchEmployeeByEmailAndNotID(ctx context.Context, id uuid.UUID, email string) (dtos.EmployeeResponse, error) {
	ret := _m.Called(ctx, id, email)

	if len(ret) == 0 {
		panic("no return value specified for SearchEmployeeByEmailAndNotID")
	}

	var r0 dtos.EmployeeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) (dtos.EmployeeResponse, error)); ok {
		return rf(ctx, id, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) dtos.EmployeeResponse); ok {
		r0 = rf(ctx, id, email)
	} else {
		r0 = ret.Get(0).(dtos.EmployeeResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, id, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchEmployeeByID provides a mock function with given fields: ctx, id
func (_m *Employee) SearchEmployeeByID(ctx context.Context, id uuid.UUID) (dtos.EmployeeResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for SearchEmployeeByID")
	}

	var r0 dtos.EmployeeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (dtos.EmployeeResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) dtos.EmployeeResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(dtos.EmployeeResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEmployee provides a mock function with given fields: ctx, empl, id
func (_m *Employee) UpdateEmployee(ctx context.Context, empl models.Employee, id uuid.UUID) error {
	ret := _m.Called(ctx, empl, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Employee, uuid.UUID) error); ok {
		r0 = rf(ctx, empl, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
