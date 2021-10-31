package cart_test

// import (
// 	"altaStore/business"
// 	"altaStore/business/cart"
// 	cartMock "altaStore/business/cart/mocks"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// const (
// 	userID    = 1
// 	status    = "active"
// 	productID = 1
// 	price     = 25000
// 	qty       = 5
// 	do        = "addition"
// )

// var (
// 	cartService       cart.Service
// 	cartRepository    cartMock.Repository
// 	productRepository productMock.Repository
// 	cartData          cart.Cart
// 	cartDetailData    cart.CartDetail
// 	addToCartData     cart.AddToCartSpec
// )

// func TestMain(m *testing.M) {
// 	setup()
// 	os.Exit(m.Run())
// }

// func TestAddToCart(t *testing.T) {
// 	t.Run("Expect add to cart success", func(t *testing.T) {
// 		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
// 		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
// 		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
// 		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(nil, business.ErrInternalServerError).Once()
// 		cartRepository.On("InsertCartDetail", mock.AnythingOfType("cart.CartDetail")).Return(nil).Once()

// 		err := cartService.AddToCart(addToCartData)

// 		assert.Nil(t, err)

// 	})

// }

// func setup() {

// 	cartData = cart.NewCart(
// 		userID,
// 		status,
// 		time.Now(),
// 	)

// 	addToCartData = cart.AddToCartSpec{
// 		UserID:    userID,
// 		ProductID: productID,
// 		Price:     price,
// 		Quantity:  qty,
// 		Do:        do,
// 	}

// 	cartService = cart.NewService(&cartRepository, &productRepository)
// }
