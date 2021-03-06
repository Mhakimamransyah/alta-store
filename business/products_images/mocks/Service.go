// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	productsimages "altaStore/business/products_images"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// InsertNewImages provides a mock function with given fields: products_image, files, createdById
func (_m *Service) InsertNewImages(products_image *productsimages.ProductImages, files []*multipart.FileHeader, createdById int) error {
	ret := _m.Called(products_image, files, createdById)

	var r0 error
	if rf, ok := ret.Get(0).(func(*productsimages.ProductImages, []*multipart.FileHeader, int) error); ok {
		r0 = rf(products_image, files, createdById)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
