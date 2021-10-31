// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	cart "altaStore/business/cart"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateCart provides a mock function with given fields: _a0
func (_m *Repository) CreateCart(_a0 cart.Cart) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(cart.Cart) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindProductOnCartDetail provides a mock function with given fields: cartID, productID
func (_m *Repository) FindProductOnCartDetail(cartID uint, productID uint) (*cart.CartDetail, error) {
	ret := _m.Called(cartID, productID)

	var r0 *cart.CartDetail
	if rf, ok := ret.Get(0).(func(uint, uint) *cart.CartDetail); ok {
		r0 = rf(cartID, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.CartDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(cartID, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveCart provides a mock function with given fields: userID
func (_m *Repository) GetActiveCart(userID uint) (*cart.Cart, error) {
	ret := _m.Called(userID)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(uint) *cart.Cart); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCartDetailByCartID provides a mock function with given fields: cartID
func (_m *Repository) GetCartDetailByCartID(cartID uint) ([]cart.CartDetail, error) {
	ret := _m.Called(cartID)

	var r0 []cart.CartDetail
	if rf, ok := ret.Get(0).(func(uint) []cart.CartDetail); ok {
		r0 = rf(cartID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.CartDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCartDetail provides a mock function with given fields: cartDetail
func (_m *Repository) InsertCartDetail(cartDetail cart.CartDetail) error {
	ret := _m.Called(cartDetail)

	var r0 error
	if rf, ok := ret.Get(0).(func(cart.CartDetail) error); ok {
		r0 = rf(cartDetail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateQuantity provides a mock function with given fields: cartID, productID, qty
func (_m *Repository) UpdateQuantity(cartID uint, productID uint, qty uint) error {
	ret := _m.Called(cartID, productID, qty)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint) error); ok {
		r0 = rf(cartID, productID, qty)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatusCart provides a mock function with given fields: cartID, status
func (_m *Repository) UpdateStatusCart(cartID uint, status string) error {
	ret := _m.Called(cartID, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string) error); ok {
		r0 = rf(cartID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
