package address_test

import (
	"altaStore/business"
	"altaStore/business/address"
	addressMock "altaStore/business/address/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userID      uint = 1
	name             = "Receiver Name"
	phoneNumber      = "083213216545"
	street           = "Street"
	city             = "City"
	province         = "Province"
	district         = "District"
	postalCode       = 11111
)

var (
	addressService address.Service
	addressRepo    addressMock.Repository

	insertAddressData  address.InsertAddressSpec
	insertAddressData1 address.InsertAddressSpec
	fakeAddressData    address.InsertAddressSpec
	newAddress         address.Address
	listAddress        []address.Address
	emptyListAddress   []address.Address
	addressType        string
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestInsertAddress(t *testing.T) {

	t.Run("Expect insert address success", func(t *testing.T) {
		addressRepo.On("GetDefaultAddress", mock.AnythingOfType("uint")).Return(nil, business.ErrNotFound).Once()
		addressRepo.On("InsertAddress", mock.AnythingOfType("address.Address")).Return(nil).Once()

		err := addressService.InsertAddress(insertAddressData)
		assert.Nil(t, err)
	})

	t.Run("Expect insert address fail and return ErrInvalid", func(t *testing.T) {

		err := addressService.InsertAddress(fakeAddressData)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInvalidSpec)
	})

	t.Run("Expect insert address failed and return ErrInternalServerError", func(t *testing.T) {
		addressRepo.On("GetDefaultAddress", mock.AnythingOfType("uint")).Return(&newAddress, nil).Once()
		addressRepo.On("UpdateDefaultAddress", mock.AnythingOfType("address.Address")).Return(business.ErrInternalServerError).Once()

		err := addressService.InsertAddress(insertAddressData1)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect insert address fail and return ErrInternalServerError", func(t *testing.T) {
		addressRepo.On("GetDefaultAddress", mock.AnythingOfType("uint")).Return(nil, business.ErrNotFound).Once()
		addressRepo.On("InsertAddress", mock.AnythingOfType("address.Address")).Return(business.ErrInternalServerError).Once()

		err := addressService.InsertAddress(insertAddressData)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

}

func TestGetDefaultAddress(t *testing.T) {
	t.Run("Expect return Address and nil of error", func(t *testing.T) {
		addressRepo.On("GetDefaultAddress", mock.AnythingOfType("uint")).Return(&newAddress, nil).Once()

		address, err := addressService.GetDefaultAddress(userID)
		assert.Nil(t, err)
		assert.NotNil(t, address)
		assert.Equal(t, address, &newAddress)
	})

	t.Run("Expect return nil and error", func(t *testing.T) {
		addressRepo.On("GetDefaultAddress", mock.AnythingOfType("uint")).Return(nil, business.ErrInternalServerError).Once()

		address, err := addressService.GetDefaultAddress(userID)
		assert.Nil(t, address)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

}

func TestGetAllAddress(t *testing.T) {
	t.Run("Expect return array of Address and nil of error", func(t *testing.T) {
		addressRepo.On("GetAllAddress", mock.AnythingOfType("uint")).Return(listAddress, nil).Once()

		address, err := addressService.GetAllAddress(userID)

		assert.Nil(t, err)
		assert.NotNil(t, address)
		assert.Equal(t, address, listAddress)
	})

	t.Run("Expect return nil and error", func(t *testing.T) {
		addressRepo.On("GetAllAddress", mock.AnythingOfType("uint")).Return(emptyListAddress, business.ErrInternalServerError).Once()

		address, err := addressService.GetAllAddress(userID)
		assert.NotNil(t, err)
		assert.Equal(t, address, emptyListAddress)
		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func setup() {
	addressType = ""
	insertAddressData = address.InsertAddressSpec{
		UserID:      userID,
		Name:        name,
		PhoneNumber: phoneNumber,
		Street:      street,
		City:        city,
		Province:    province,
		District:    district,
		PostalCode:  postalCode,
		AddressType: &addressType,
		IsDefault:   false,
	}

	insertAddressData1 = insertAddressData
	insertAddressData1.IsDefault = true

	fakeAddressData = address.InsertAddressSpec{
		UserID:      userID,
		Name:        name,
		PhoneNumber: "asdkajshdlkasdhl",
		Street:      street,
		City:        city,
		Province:    province,
		District:    district,
		PostalCode:  postalCode,
		AddressType: &addressType,
		IsDefault:   false,
	}

	newAddress = address.Address{
		ID:          1,
		UserID:      userID,
		Name:        name,
		PhoneNumber: phoneNumber,
		Street:      street,
		City:        city,
		Province:    province,
		District:    district,
		PostalCode:  postalCode,
		AddressType: &addressType,
		IsDefault:   true,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   &time.Time{},
	}

	emptyListAddress = []address.Address{}
	listAddress = append(listAddress, newAddress)
	addressService = address.NewService(&addressRepo)
}
