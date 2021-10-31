// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	products "altaStore/business/products"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateProducts provides a mock function with given fields: _a0, createdById
func (_m *Repository) CreateProducts(_a0 *products.Products, createdById int) (*products.Products, error) {
	ret := _m.Called(_a0, createdById)

	var r0 *products.Products
	if rf, ok := ret.Get(0).(func(*products.Products, int) *products.Products); ok {
		r0 = rf(_a0, createdById)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.Products)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.Products, int) error); ok {
		r1 = rf(_a0, createdById)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProducts provides a mock function with given fields: _a0
func (_m *Repository) DeleteProducts(_a0 *products.Products) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*products.Products) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProducts provides a mock function with given fields: filter
func (_m *Repository) GetAllProducts(filter products.FilterProducts) (*[]products.Products, error) {
	ret := _m.Called(filter)

	var r0 *[]products.Products
	if rf, ok := ret.Get(0).(func(products.FilterProducts) *[]products.Products); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]products.Products)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(products.FilterProducts) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailProducts provides a mock function with given fields: id_products
func (_m *Repository) GetDetailProducts(id_products int) (*products.Products, error) {
	ret := _m.Called(id_products)

	var r0 *products.Products
	if rf, ok := ret.Get(0).(func(int) *products.Products); ok {
		r0 = rf(id_products)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.Products)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id_products)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReduceStocks provides a mock function with given fields: cost_reduce, id_products
func (_m *Repository) ReduceStocks(cost_reduce int, id_products int) error {
	ret := _m.Called(cost_reduce, id_products)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(cost_reduce, id_products)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProducts provides a mock function with given fields: id_products, _a1, modifiedById
func (_m *Repository) UpdateProducts(id_products int, _a1 *products.Products, modifiedById int) error {
	ret := _m.Called(id_products, _a1, modifiedById)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *products.Products, int) error); ok {
		r0 = rf(id_products, _a1, modifiedById)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
