// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	cart "altaStore/business/cart"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddToCart provides a mock function with given fields: addToCartSpec
func (_m *Service) AddToCart(addToCartSpec cart.AddToCartSpec) error {
	ret := _m.Called(addToCartSpec)

	var r0 error
	if rf, ok := ret.Get(0).(func(cart.AddToCartSpec) error); ok {
		r0 = rf(addToCartSpec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetActiveCart provides a mock function with given fields: userID
func (_m *Service) GetActiveCart(userID uint) (cart.ActiveCart, error) {
	ret := _m.Called(userID)

	var r0 cart.ActiveCart
	if rf, ok := ret.Get(0).(func(uint) cart.ActiveCart); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(cart.ActiveCart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
