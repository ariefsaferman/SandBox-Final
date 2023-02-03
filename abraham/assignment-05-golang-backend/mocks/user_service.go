// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dtos "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	mock "github.com/stretchr/testify/mock"

	models "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetUsers provides a mock function with given fields:
func (_m *UserService) GetUsers() ([]*models.User, error) {
	ret := _m.Called()

	var r0 []*models.User
	if rf, ok := ret.Get(0).(func() []*models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: loginRequest
func (_m *UserService) LoginUser(loginRequest dtos.LoginRequest) (*dtos.TokenResponse, error) {
	ret := _m.Called(loginRequest)

	var r0 *dtos.TokenResponse
	if rf, ok := ret.Get(0).(func(dtos.LoginRequest) *dtos.TokenResponse); ok {
		r0 = rf(loginRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.TokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dtos.LoginRequest) error); ok {
		r1 = rf(loginRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: registerRequest
func (_m *UserService) RegisterUser(registerRequest dtos.RegisterRequest) (*dtos.RegisterResponse, error) {
	ret := _m.Called(registerRequest)

	var r0 *dtos.RegisterResponse
	if rf, ok := ret.Get(0).(func(dtos.RegisterRequest) *dtos.RegisterResponse); ok {
		r0 = rf(registerRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.RegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dtos.RegisterRequest) error); ok {
		r1 = rf(registerRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
