// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	admins "altaStore/business/admins"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// FindAdminById provides a mock function with given fields: id_admins
func (_m *Service) FindAdminById(id_admins int) (*admins.Admins, error) {
	ret := _m.Called(id_admins)

	var r0 *admins.Admins
	if rf, ok := ret.Get(0).(func(int) *admins.Admins); ok {
		r0 = rf(id_admins)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admins.Admins)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id_admins)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAdminByUsername provides a mock function with given fields: username
func (_m *Service) FindAdminByUsername(username string) (*admins.Admins, error) {
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

// FindAllAdmin provides a mock function with given fields: offset, limit
func (_m *Service) FindAllAdmin(offset int, limit int) (*[]admins.Admins, error) {
	ret := _m.Called(offset, limit)

	var r0 *[]admins.Admins
	if rf, ok := ret.Get(0).(func(int, int) *[]admins.Admins); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]admins.Admins)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertAdmin provides a mock function with given fields: admin_spec, createdById
func (_m *Service) InsertAdmin(admin_spec admins.AdminSpec, createdById int) error {
	ret := _m.Called(admin_spec, createdById)

	var r0 error
	if rf, ok := ret.Get(0).(func(admins.AdminSpec, int) error); ok {
		r0 = rf(admin_spec, createdById)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginAdmin provides a mock function with given fields: username, password
func (_m *Service) LoginAdmin(username string, password string) (*admins.Admins, error) {
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

// ModifyAdmin provides a mock function with given fields: username, admin, modifiedById
func (_m *Service) ModifyAdmin(username string, admin admins.AdminUpdatable, modifiedById int) error {
	ret := _m.Called(username, admin, modifiedById)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, admins.AdminUpdatable, int) error); ok {
		r0 = rf(username, admin, modifiedById)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
