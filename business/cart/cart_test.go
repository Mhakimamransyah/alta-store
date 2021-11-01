package cart_test

import (
	"altaStore/business"
	"altaStore/business/cart"
	cartMock "altaStore/business/cart/mocks"
	"altaStore/business/products"
	productMock "altaStore/business/products/mocks"
	productsimages "altaStore/business/products_images"

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

	//produck const
	productIDProduct   = 1
	stock              = 100
	productName        = "product1"
	productPrice       = 100000
	productDescription = "product description"
	weight             = 1.0
	productStatus      = "active"
	adminID            = 1
	categoriesID       = 1
)

var (
	cartService        cart.Service
	cartRepository     cartMock.Repository
	productRepository  productMock.Repository
	cartData           cart.Cart
	cartDetailData     *cart.CartDetail
	addToCartData      cart.AddToCartSpec
	addToCartData1     cart.AddToCartSpec
	addToCartData2     cart.AddToCartSpec
	addToCartData3     cart.AddToCartSpec
	fakeAddToCartData  cart.AddToCartSpec
	prod               *products.Products
	activeCart         *cart.Cart
	responeActiveCart  cart.ActiveCart
	activeCartResponse cart.ActiveCart
	listCartDetail     []cart.CartDetail
	cartDetail1        cart.CartDetail
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestAddToCart(t *testing.T) {
	t.Run("Expect fail on add to cart and return ErrInvalidSpec", func(t *testing.T) {
		err := cartService.AddToCart(fakeAddToCartData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInvalidSpec)

	})

	t.Run("Expect fail add to cart and return ErrInternalServiceError", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)

	})

	t.Run("Expect fail add to cart and return ErrProductNotFound", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(nil, business.ErrProductNotFound).Once()

		err := cartService.AddToCart(addToCartData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrProductNotFound)

	})

	t.Run("Expect fail add to cart and return ErrProductOOS", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()

		err := cartService.AddToCart(addToCartData1)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrProductOOS)

	})

	t.Run("Expect fail add to cart and return ErrInternalServerError", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData2)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)

	})

	t.Run("Expect fail add to cart and return ErrInternalServerError", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData2)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)

	})

	t.Run("Expect success remove product from cartdetail", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil).Once()

		err := cartService.AddToCart(addToCartData2)
		assert.Nil(t, err)
	})

	t.Run("Expect fail add to cart return ErrInternalServerError", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect fail add to cart return ErrInternalServerError", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect success update cartdetails and nil of error", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(nil).Once()

		err := cartService.AddToCart(addToCartData)
		assert.Nil(t, err)
	})

	t.Run("Expect fail add to cart return ErrInternalServerError", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData3)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect fail add to cart return ErrInternalServerError", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData3)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect success subtract quantity from cartdetail nil of error", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, nil).Once()
		cartRepository.On("UpdateQuantity", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil).Once()

		err := cartService.AddToCart(addToCartData3)
		assert.Nil(t, err)
	})

	t.Run("Expect fail add to cart return ErrInternalServerError ", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, business.ErrInternalServerError).Once()
		cartRepository.On("InsertCartDetail", mock.AnythingOfType("cart.CartDetail")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect fail add to cart return ErrInternalServerError ", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, business.ErrInternalServerError).Once()
		cartRepository.On("InsertCartDetail", mock.AnythingOfType("cart.CartDetail")).Return(nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServerError)
	})

	t.Run("Expect success add to cart and nil of error ", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, business.ErrInternalServerError).Once()
		cartRepository.On("InsertCartDetail", mock.AnythingOfType("cart.CartDetail")).Return(nil).Once()
		productRepository.On("UpdateStocks", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil).Once()

		err := cartService.AddToCart(addToCartData)
		assert.Nil(t, err)
	})

	t.Run("Expect fail to subtraction because product not found", func(t *testing.T) {
		cartRepository.On("GetActiveCart", uint(userID)).Return(nil, business.ErrInternalServerError).Once()
		cartRepository.On("CreateCart", mock.AnythingOfType("cart.Cart")).Return(nil).Once()
		cartRepository.On("GetActiveCart", uint(userID)).Return(&cartData, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()
		cartRepository.On("FindProductOnCartDetail", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(cartDetailData, business.ErrInternalServerError).Once()

		err := cartService.AddToCart(addToCartData3)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrAddToCart)
	})
}

func TestGetActiveCart(t *testing.T) {
	t.Run("Expect return empty ActiveCart struct and nil of error", func(t *testing.T) {
		cartRepository.On("GetActiveCart", mock.AnythingOfType("uint")).Return(nil, business.ErrNotFound).Once()

		activeCart, err := cartService.GetActiveCart(userID)

		assert.Nil(t, err)
		assert.Equal(t, activeCart, responeActiveCart)

	})

	t.Run("Expect return ActiveCart struct and nil of error", func(t *testing.T) {
		cartRepository.On("GetActiveCart", mock.AnythingOfType("uint")).Return(activeCart, nil).Once()
		cartRepository.On("GetCartDetailByCartID", mock.AnythingOfType("uint")).Return(nil, business.ErrNotFound).Once()

		activeCart, err := cartService.GetActiveCart(userID)
		assert.Nil(t, err)
		assert.NotNil(t, activeCart)

	})

	t.Run("Expect return ActiveCart struct and nil of error", func(t *testing.T) {
		cartRepository.On("GetActiveCart", mock.AnythingOfType("uint")).Return(activeCart, nil).Once()
		cartRepository.On("GetCartDetailByCartID", mock.AnythingOfType("uint")).Return(listCartDetail, nil).Once()
		productRepository.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()

		activeCart, err := cartService.GetActiveCart(userID)
		assert.Nil(t, err)
		assert.NotNil(t, activeCart)

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

	addToCartData1 = addToCartData
	addToCartData1.Quantity = 200

	addToCartData2 = addToCartData
	addToCartData2.Quantity = 0

	addToCartData3 = addToCartData
	addToCartData3.Do = "subtraction"

	fakeAddToCartData = cart.AddToCartSpec{
		UserID:    userID,
		ProductID: productID,
		Price:     price,
		Quantity:  qty,
		Do:        "",
	}

	prod = &products.Products{
		ID:              productIDProduct,
		Stock:           stock,
		Title:           productName,
		Price:           productPrice,
		Description:     productDescription,
		Weight:          weight,
		Status:          productStatus,
		AdminID:         adminID,
		CategoriesID:    categoriesID,
		Created_at:      time.Time{},
		Updated_at:      time.Time{},
		Deleted_at:      time.Time{},
		Products_images: []productsimages.ProductImages{},
	}

	cartDetailData = &cart.CartDetail{
		ID:        1,
		CartID:    1,
		ProductID: productID,
		Price:     productPrice,
		Quantity:  qty,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}

	activeCart = &cart.Cart{
		ID:        1,
		UserID:    userID,
		Status:    "active",
		AddressID: new(uint),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}

	responeActiveCart = cart.ActiveCart{}
	activeCartResponse = cart.ActiveCart{
		CartID:      1,
		Status:      "active",
		AddressID:   new(uint),
		CartDetails: []cart.ActiveCartDetail{},
	}

	cartDetail1 = cart.CartDetail{
		ID:        1,
		CartID:    1,
		ProductID: 1,
		Price:     price,
		Quantity:  qty,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}

	listCartDetail = append(listCartDetail, cartDetail1)

	cartService = cart.NewService(&cartRepository, &productRepository)
}
