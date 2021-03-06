// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	admins "altaStore/business/admins"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateAdmin provides a mock function with given fields: admin
func (_m *Repository) CreateAdmin(admin *admins.Admins) error {
	ret := _m.Called(admin)

	var r0 error
	if rf, ok := ret.Get(0).(func(*admins.Admins) error); ok {
		r0 = rf(admin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdmin provides a mock function with given fields: limit, offset
func (_m *Repository) GetAdmin(limit int, offset int) (*[]admins.Admins, error) {
	ret := _m.Called(limit, offset)

	var r0 *[]admins.Admins
	if rf, ok := ret.Get(0).(func(int, int) *[]admins.Admins); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]admins.Admins)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdminById provides a mock function with given fields: id
func (_m *Repository) GetAdminById(id int) (*admins.Admins, error) {
	ret := _m.Called(id)

	var r0 *admins.Admins
	if rf, ok := ret.Get(0).(func(int) *admins.Admins); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admins.Admins)
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

// GetAdminByUsername provides a mock function with given fields: username
func (_m *Repository) GetAdminByUsername(username string) (*admins.Admins, error) {
	ret := _m.Called(username)

	var r0 *admins.Admins
	if rf, ok := ret.Get(0).(func(string) *admins.Admins); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admins.Admins)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginAdmin provides a mock function with given fields: username, password
func (_m *Repository) LoginAdmin(username string, password string) (*admins.Admins, error) {
	ret := _m.Called(username, password)

	var r0 *admins.Admins
	if rf, ok := ret.Get(0).(func(string, string) *admins.Admins); ok {
		r0 = rf(username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admins.Admins)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAdmin provides a mock function with given fields: admin
func (_m *Repository) UpdateAdmin(admin *admins.Admins) error {
	ret := _m.Called(admin)

	var r0 error
	if rf, ok := ret.Get(0).(func(*admins.Admins) error); ok {
		r0 = rf(admin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
