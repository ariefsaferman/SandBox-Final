// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// GetTransactions provides a mock function with given fields: _a0, _a1
func (_m *TransactionRepository) GetTransactions(_a0 uint, _a1 entity.TransactionParams) ([]*entity.Transaction, int64, int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*entity.Transaction
	if rf, ok := ret.Get(0).(func(uint, entity.TransactionParams) []*entity.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Transaction)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(uint, entity.TransactionParams) int64); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(uint, entity.TransactionParams) int); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(uint, entity.TransactionParams) error); ok {
		r3 = rf(_a0, _a1)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// TopUp provides a mock function with given fields: req
func (_m *TransactionRepository) TopUp(req entity.Transaction) (*entity.Transaction, error) {
	ret := _m.Called(req)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(entity.Transaction) *entity.Transaction); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Transaction) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: req
func (_m *TransactionRepository) Transfer(req entity.Transaction) (*entity.Transaction, error) {
	ret := _m.Called(req)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(entity.Transaction) *entity.Transaction); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Transaction) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}