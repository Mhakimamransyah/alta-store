package user_test

import (
	"altaStore/business"
	"altaStore/business/user"
	userMock "altaStore/business/user/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id          = 1
	name        = "name"
	email       = "email@example.com"
	phoneNumber = "08123123123"
	password    = "password"
	skip        = 0
	rowPerPage  = 10
)

var (
	userService        user.Service
	userRepo           userMock.Repository
	utilService        userMock.Util
	userData           user.User
	insertUserData     user.InsertUserSpec
	fakeinsertUserData user.InsertUserSpec
	usersData          []user.User
	listUser           []user.User
	hashPassword1      []byte
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindUserByID(t *testing.T) {
	t.Run("Expected found the user", func(t *testing.T) {
		userRepo.On("FindUserByID", mock.AnythingOfType("int")).Return(&userData, nil).Once()

		user, err := userService.FindUserByID(id)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, phoneNumber, user.PhoneNumber)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expected user not found", func(t *testing.T) {
		userRepo.On("FindUserByID", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		user, err := userService.FindUserByID(id)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("Expected found the user", func(t *testing.T) {
		userRepo.On("FindUserByEmail", mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := userService.FindUserByEmail(email)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, phoneNumber, user.PhoneNumber)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expected user not found", func(t *testing.T) {
		userRepo.On("FindUserByEmail", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		user, err := userService.FindUserByEmail(email)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestFindAllUser(t *testing.T) {
	t.Run("Expected found the user", func(t *testing.T) {
		userRepo.On("FindAllUser", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(listUser, business.ErrNotFound).Once()

		_, err := userService.FindAllUser(skip, rowPerPage)

		assert.NotNil(t, err)

	})

	t.Run("Expected return empty array", func(t *testing.T) {
		userRepo.On("FindAllUser", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(usersData, nil).Once()

		users, err := userService.FindAllUser(skip, rowPerPage)

		assert.Nil(t, err)
		assert.Equal(t, usersData, users)
	})
}

func TestInsertUser(t *testing.T) {

	t.Run("Expected insert user fail and return ErrInvalidSpec", func(t *testing.T) {

		err := userService.InsertUser(fakeinsertUserData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInvalidSpec)
	})

	t.Run("Expected insert user success", func(t *testing.T) {
		utilService.On("EncryptPassword", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServerError).Once()

		// userRepo.On("InsertUser", mock.AnythingOfType("user.User")).Return(nil).Once()

		err := userService.InsertUser(insertUserData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expected insert user success", func(t *testing.T) {
		utilService.On("EncryptPassword", mock.AnythingOfType("string")).Return(hashPassword1, nil).Once()
		userRepo.On("InsertUser", mock.AnythingOfType("user.User")).Return(business.ErrRegister).Once()

		err := userService.InsertUser(insertUserData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrRegister)
	})

	t.Run("Expected insert user success", func(t *testing.T) {
		utilService.On("EncryptPassword", mock.AnythingOfType("string")).Return(hashPassword1, nil).Once()
		userRepo.On("InsertUser", mock.AnythingOfType("user.User")).Return(nil).Once()

		err := userService.InsertUser(insertUserData)

		assert.Nil(t, err)
	})

}

func setup() {

	userData = user.NewUser(
		name,
		email,
		phoneNumber,
		password,
		time.Now(),
	)

	listUser = append(listUser, userData)

	insertUserData = user.InsertUserSpec{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}

	fakeinsertUserData = user.InsertUserSpec{
		Name:        "",
		Email:       "",
		PhoneNumber: phoneNumber,
		Password:    password,
	}

	userService = user.NewService(&userRepo, &utilService)
}
