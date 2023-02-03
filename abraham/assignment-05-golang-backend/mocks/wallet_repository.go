// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	mock "github.com/stretchr/testify/mock"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// CreateWallet provides a mock function with given fields: id
func (_m *WalletRepository) CreateWallet(id int) (*models.Wallet, error) {
	ret := _m.Called(id)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(int) *models.Wallet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserIDByWalletNumber provides a mock function with given fields: walletNumber
func (_m *WalletRepository) GetUserIDByWalletNumber(walletNumber int) (int, error) {
	ret := _m.Called(walletNumber)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(walletNumber)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(walletNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletBalance provides a mock function with given fields: id
func (_m *WalletRepository) GetWalletBalance(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletNumber provides a mock function with given fields: id
func (_m *WalletRepository) GetWalletNumber(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryWallet provides a mock function with given fields: id
func (_m *WalletRepository) QueryWallet(id int) (*models.Wallet, error) {
	ret := _m.Called(id)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(int) *models.Wallet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWalletBalance provides a mock function with given fields: amount, id
func (_m *WalletRepository) UpdateWalletBalance(amount int, id int) (*models.Wallet, error) {
	ret := _m.Called(amount, id)

	var r0 *models.Wallet
	if rf, ok := ret.Get(0).(func(int, int) *models.Wallet); ok {
		r0 = rf(amount, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(amount, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWalletRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletRepository creates a new instance of WalletRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletRepository(t mockConstructorTestingTNewWalletRepository) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}