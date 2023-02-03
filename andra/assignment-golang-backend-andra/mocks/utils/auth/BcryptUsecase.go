// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	entity "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// BcryptUsecase is an autogenerated mock type for the BcryptUsecase type
type BcryptUsecase struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: hashedPwd, inputPwd
func (_m *BcryptUsecase) ComparePassword(hashedPwd string, inputPwd string) bool {
	ret := _m.Called(hashedPwd, inputPwd)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hashedPwd, inputPwd)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GenerateAccessToken provides a mock function with given fields: req
func (_m *BcryptUsecase) GenerateAccessToken(req entity.User) dto.LoginResponse {
	ret := _m.Called(req)

	var r0 dto.LoginResponse
	if rf, ok := ret.Get(0).(func(entity.User) dto.LoginResponse); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(dto.LoginResponse)
	}

	return r0
}

// HashAndSalt provides a mock function with given fields: pwd
func (_m *BcryptUsecase) HashAndSalt(pwd string) string {
	ret := _m.Called(pwd)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(pwd)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewBcryptUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewBcryptUsecase creates a new instance of BcryptUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBcryptUsecase(t mockConstructorTestingTNewBcryptUsecase) *BcryptUsecase {
	mock := &BcryptUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
