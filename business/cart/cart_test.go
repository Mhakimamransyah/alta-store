package cart_test

import (
	"altaStore/business/cart"
	cartMock "altaStore/business/cart/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userID    = 1
	status    = "active"
	productID = 1
	price     = 25000
	qty       = 5
	do        = "addition"
)

var (
	cartService    cart.Service
	cartRepository cartMock.Repository

	cartData      cart.Cart
	addToCartData cart.AddToCartSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestAddToCart(t *testing.T) {
	t.Run("Expect add to cart success", func(t *testing.T) {
		cartRepository.On("AddToCart", mock.AnythingOfType("cart.AddToCartSpec")).Return(nil).Once()

		err := cartService.AddToCart(addToCartData)

		assert.Nil(t, err)

	})
}

func setup() {

	cartData = cart.NewCart(
		userID,
		status,
		time.Now(),
	)

	addToCartData = cart.AddToCartSpec{
		UserID:    userID,
		ProductID: productID,
		Price:     price,
		Quantity:  qty,
		Do:        do,
	}

	cartService = cart.NewService(&cartRepository)
}
