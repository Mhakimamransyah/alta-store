package transaction

import (
	"altaStore/business"
	"altaStore/business/address"
	"altaStore/business/cart"
	"altaStore/business/products"
	"altaStore/util/validator"
	"strconv"
	"time"
)

const (
	RekBank   = "Bank BCA"
	RekNumber = "0321555486548"
	RekName   = "Alta Store"
)

type CheckoutSpec struct {
	CartID      uint `validate:"gt=0"`
	UserID      uint `validate:"gt=0"`
	ShippingFee uint `validate:"gte=0"`
	AddressID   uint `validate:"gte=0"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository        Repository
	cartRepository    cart.Repository
	addressRepository address.Repository
	productRepository products.Repository
}

func NewService(
	repository Repository,
	cartRepository cart.Repository,
	addressRepo address.Repository,
	productRepo products.Repository) Service {
	return &service{
		repository,
		cartRepository,
		addressRepo,
		productRepo,
	}
}

func (s *service) Checkout(checkoutSpec CheckoutSpec) (*CheckoutResponse, error) {
	err := validator.GetValidator().Struct(checkoutSpec)
	if err != nil {
		return nil, business.ErrInvalidSpec
	}

	getActiveCart, err := s.cartRepository.GetActiveCart(checkoutSpec.UserID)

	if err != nil || getActiveCart.ID != checkoutSpec.CartID {
		return nil, business.ErrActiveCartNotFound
	}

	cartDetails, err := s.cartRepository.GetCartDetailByCartID(checkoutSpec.CartID)

	if len(cartDetails) == 0 || err != nil {
		return nil, business.ErrCartDetailEmpty
	}

	addressTransaction, err := s.addressRepository.GetAddressForTransaction(checkoutSpec.UserID, checkoutSpec.AddressID)

	if err != nil {
		return nil, business.ErrAddressNotFound
	}

	newAddress := ToAddressResponse(*addressTransaction)

	var product *products.Products
	var transactionProductList []Product
	var totalPrice uint

	for _, value := range cartDetails {
		product, _ = s.productRepository.GetDetailProducts(int(value.ProductID))
		transactionProductList = append(transactionProductList, ToTransactionProduct(*product, value.Price, value.Quantity))
		totalPrice += (value.Price * value.Quantity)
	}

	var invoiceNumber string
	t := time.Now()
	year := strconv.Itoa(t.Year())
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	invoiceNumber = "INV/" + year + month + day + "/" + strconv.Itoa(int(checkoutSpec.UserID)) + "/" + strconv.Itoa(int(checkoutSpec.CartID))

	transactionData := NewTransaction(
		checkoutSpec.CartID,
		invoiceNumber,
		"waiting",
		totalPrice,
		checkoutSpec.ShippingFee,
		t,
	)

	err = s.repository.CreateTransaction(transactionData)
	if err != nil {
		return nil, err
	}

	err = s.cartRepository.UpdateStatusCart(checkoutSpec.CartID, "checkout")
	if err != nil {
		return nil, err
	}

	err = s.cartRepository.UpdateAddressID(checkoutSpec.CartID, checkoutSpec.AddressID)
	if err != nil {
		return nil, err
	}

	response := ToCheckoutResponse(
		invoiceNumber,
		RekBank,
		RekNumber,
		RekName,
		t,
		totalPrice,
		checkoutSpec.ShippingFee,
		(totalPrice + checkoutSpec.ShippingFee),
		newAddress,
		transactionProductList,
	)
	return &response, nil
}
