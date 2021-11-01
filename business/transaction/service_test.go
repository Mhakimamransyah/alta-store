package transaction_test

import (
	"altaStore/business"
	"altaStore/business/address"
	addressMock "altaStore/business/address/mocks"
	"altaStore/business/cart"
	cartMock "altaStore/business/cart/mocks"
	"altaStore/business/products"
	productMock "altaStore/business/products/mocks"
	productsimages "altaStore/business/products_images"
	"altaStore/business/transaction"
	trxMock "altaStore/business/transaction/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	cartID uint = 1
	userID uint = 1
)

var (
	trxService        transaction.Service
	trxServiceMock    trxMock.Service
	trxRepo           trxMock.Repository
	cartRepo          cartMock.Repository
	addressRepo       addressMock.Repository
	productRepo       productMock.Repository
	carts             []cart.Cart
	emptycarts        []cart.Cart
	cart1             cart.Cart
	checkoutResponse  []*transaction.CheckoutResponse
	transactions      []transaction.Transaction
	trx1              transaction.Transaction
	trx2              *transaction.Transaction
	checkoutResponse1 *transaction.CheckoutResponse
	cart2             *cart.Cart
	cart3             *cart.Cart
	address1          *address.TransactionAddress
	cartDetails       []cart.CartDetail
	cartDetail        cart.CartDetail
	prod              *products.Products
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindTransactionByInvoice(t *testing.T) {
	t.Run("Expect return nil transaction and ErrTransactionNotFound  ", func(t *testing.T) {
		trxRepo.On("FindTransactionByInvoice", mock.AnythingOfType("string")).Return(nil, business.ErrTransactionNotFound).Once()

		transac, err := trxService.FindTransactionByInvoice("asdasdasd", userID)

		assert.Nil(t, transac)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrTransactionNotFound)
	})

	t.Run("Expect errTransaction Access ", func(t *testing.T) {
		trxRepo.On("FindTransactionByInvoice", mock.AnythingOfType("string")).Return(trx2, nil).Once()
		cartRepo.On("FindCartByID", mock.AnythingOfType("uint")).Return(cart2, business.ErrTransactionAccess).Once()

		transac, err := trxService.FindTransactionByInvoice("asdasdasd", userID)

		assert.Nil(t, transac)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrTransactionAccess)
	})

	t.Run("Expect find transaction ", func(t *testing.T) {
		trxRepo.On("FindTransactionByInvoice", mock.AnythingOfType("string")).Return(trx2, nil).Once()
		cartRepo.On("FindCartByID", mock.AnythingOfType("uint")).Return(cart3, nil).Once()
		addressRepo.On("GetAddressByID", mock.AnythingOfType("uint")).Return(address1, nil).Once()
		cartRepo.On("GetCartDetailByCartID", mock.AnythingOfType("uint")).Return(cartDetails, nil).Once()
		productRepo.On("GetDetailProducts", mock.AnythingOfType("int")).Return(prod, nil).Once()

		transac, err := trxService.FindTransactionByInvoice("asdasdasd", userID)

		assert.Nil(t, err)
		assert.NotNil(t, transac)
	})
}
func TestGetAllTransaction(t *testing.T) {
	t.Run("Expect return error not found ", func(t *testing.T) {
		cartRepo.On("GetAllCartIDTransaction", mock.AnythingOfType("uint")).Return(carts, business.ErrNotFound).Once()

		allTransactions, err := trxService.GetAllTransaction(userID)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
		assert.Equal(t, allTransactions, checkoutResponse)

	})
}

func setup() {
	cart1 = cart.Cart{
		ID:        cartID,
		UserID:    userID,
		Status:    "",
		AddressID: new(uint),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}

	carts = append(carts, cart1)
	emptycarts = []cart.Cart{}
	checkoutResponse = nil

	trx1 = transaction.Transaction{
		ID:               1,
		CartID:           cartID,
		InvoiceNumber:    "INJKLAlaslkdjlalskjd",
		Status:           "",
		TotalTransaction: 0,
		ShippingFee:      0,
		PayAt:            &time.Time{},
		CancelAt:         &time.Time{},
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        &time.Time{},
	}

	transactions = append(transactions, trx1)
	trx2 = &transaction.Transaction{
		ID:               1,
		CartID:           cartID,
		InvoiceNumber:    "",
		Status:           "",
		TotalTransaction: 0,
		ShippingFee:      0,
		PayAt:            &time.Time{},
		CancelAt:         &time.Time{},
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        &time.Time{},
	}

	checkoutResponse1 = &transaction.CheckoutResponse{
		InvoiceNumber:    "",
		RekeningBank:     "",
		RekeningNumber:   "",
		RekeningName:     "",
		OrderDate:        time.Time{},
		Address:          transaction.Address{},
		Product:          []transaction.Product{},
		TotalTransaction: 0,
		ShippingFee:      0,
		TotalPay:         0,
	}

	cart2 = &cart.Cart{
		ID:        cartID,
		UserID:    3,
		Status:    "active",
		AddressID: new(uint),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}

	cart3 = &cart.Cart{
		ID:        cartID,
		UserID:    userID,
		Status:    "active",
		AddressID: new(uint),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}

	address1 = &address.TransactionAddress{
		Name:        "",
		PhoneNumber: "",
		Street:      "",
		City:        "",
		Province:    "",
		District:    "",
		PostalCode:  0,
		AddressType: new(string),
	}

	cartDetail = cart.CartDetail{
		ID:        1,
		CartID:    cartID,
		ProductID: 0,
		Price:     0,
		Quantity:  0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: &time.Time{},
	}

	cartDetails = append(cartDetails, cartDetail)

	prod = &products.Products{
		ID:              1,
		Stock:           10,
		Title:           "",
		Price:           0,
		Description:     "",
		Weight:          0,
		Status:          "",
		AdminID:         0,
		CategoriesID:    0,
		Created_at:      time.Time{},
		Updated_at:      time.Time{},
		Deleted_at:      time.Time{},
		Products_images: []productsimages.ProductImages{},
	}

	trxService = transaction.NewService(
		&trxRepo,
		&cartRepo,
		&addressRepo,
		&productRepo,
	)
}
