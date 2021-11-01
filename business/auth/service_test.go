package auth_test

import (
	"altaStore/business"
	"altaStore/business/auth"
	utilMock "altaStore/business/auth/mocks"
	"altaStore/business/user"
	userMock "altaStore/business/user/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userID      uint = 1
	name             = "name"
	email            = "email@sample.com"
	phoneNumber      = "081321324654"
	password         = "passwordsample"
	tokenJwt         = "aljksdhlaksdjlajsdla.asldkhakosdjalksdjlakjsdasd.alksdhjaklsdjklajsdkla"
)

var (
	authService    auth.Service
	userService    userMock.Service
	utilService    utilMock.UtilPassword
	userAsresponse *user.User
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestLogin(t *testing.T) {
	t.Run("Expected success login and return string token and nil of error", func(t *testing.T) {
		userService.On("FindUserByEmail", mock.AnythingOfType("string")).Return(userAsresponse, nil).Once()
		utilService.On("ComparePassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true).Once()

		token, err := authService.Login(email, password)
		assert.Nil(t, err)
		assert.NotNil(t, token)
	})

	t.Run("Expected success login and return string token and nil of error", func(t *testing.T) {
		userService.On("FindUserByEmail", mock.AnythingOfType("string")).Return(userAsresponse, nil).Once()
		utilService.On("ComparePassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(false).Once()

		token, err := authService.Login(email, password)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("Expected success login and return string token and nil of error", func(t *testing.T) {
		userService.On("FindUserByEmail", mock.AnythingOfType("string")).Return(nil, business.ErrLogin).Once()

		token, err := authService.Login(email, password)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})
}

func setup() {
	userAsresponse = &user.User{
		ID:          userID,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   &time.Time{},
	}
	authService = auth.NewService(&userService, &utilService)
}
