// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	dtos "github.com/Shepherd-Go/Back-Nlj-Internal.git/dtos"
	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"
)

// JWT is an autogenerated mock type for the JWT type
type JWT struct {
	mock.Mock
}

// PaserLoginJWT provides a mock function with given fields: value
func (_m *JWT) PaserLoginJWT(value string) (jwt.MapClaims, error) {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for PaserLoginJWT")
	}

	var r0 jwt.MapClaims
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (jwt.MapClaims, error)); ok {
		return rf(value)
	}
	if rf, ok := ret.Get(0).(func(string) jwt.MapClaims); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.MapClaims)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignedLoginToken provides a mock function with given fields: u
func (_m *JWT) SignedLoginToken(u dtos.Session) (string, error) {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for SignedLoginToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(dtos.Session) (string, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(dtos.Session) string); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(dtos.Session) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWT creates a new instance of JWT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWT(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWT {
	mock := &JWT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
