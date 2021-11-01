package transaction_test

import (
	"altaStore/business"
	addressMock "altaStore/business/address/mocks"
	"altaStore/business/cart"
	cartMock "altaStore/business/cart/mocks"
	productMock "altaStore/business/products/mocks"
	"altaStore/business/transaction"
	trxMock "altaStore/business/transaction/mocks"
	"fmt"
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
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAllTransaction(t *testing.T) {
	t.Run("Expect return error not found ", func(t *testing.T) {
		cartRepo.On("GetAllCartIDTransaction", mock.AnythingOfType("uint")).Return(carts, business.ErrNotFound).Once()

		allTransactions, err := trxService.GetAllTransaction(userID)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
		assert.Equal(t, allTransactions, checkoutResponse)

	})

	t.Run("Expect return all transaction ", func(t *testing.T) {
		cartRepo.On("GetAllCartIDTransaction", mock.AnythingOfType("uint")).Return(carts, nil).Once()
		trxRepo.On("FindAllTransaction", mock.AnythingOfType("[]uint")).Return(transactions, nil).Once()
		trxServiceMock.On("FindTransactionByInvoice", mock.AnythingOfType("string"), mock.AnythingOfType("uint")).Return(checkoutResponse1, nil).Once()

		allTransactions, err := trxService.GetAllTransaction(userID)
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println(allTransactions)
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		assert.NotNil(t, err)
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

	trxService = transaction.NewService(
		&trxRepo,
		&cartRepo,
		&addressRepo,
		&productRepo,
	)
}
