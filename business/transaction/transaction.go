package transaction

import (
	"altaStore/business/address"
	"altaStore/business/products"
	"time"
)

type Transaction struct {
	ID               uint
	CartID           uint
	InvoiceNumber    string
	Status           string
	TotalTransaction uint
	ShippingFee      uint
	PayAt            *time.Time
	CancelAt         *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

type CheckoutResponse struct {
	InvoiceNumber    string
	RekeningBank     string
	RekeningNumber   string
	RekeningName     string
	OrderDate        time.Time
	Address          Address
	Product          []Product
	TotalTransaction uint
	ShippingFee      uint
	TotalPay         uint
}

type Address struct {
	Name        string
	PhoneNumber string
	Street      string
	City        string
	Province    string
	District    string
	PostalCode  uint
	AddressType *string
}

type Product struct {
	Title             string
	Price             uint
	Quantity          uint
	TotalProductPrice uint
}

//NewTransaction create new Transaction
func NewTransaction(
	cartID uint,
	invoiceNumber string,
	status string,
	totalTransaction uint,
	shippingFee uint,
	createdAt time.Time) Transaction {

	return Transaction{
		ID:               0,
		CartID:           cartID,
		InvoiceNumber:    invoiceNumber,
		Status:           status,
		TotalTransaction: totalTransaction,
		ShippingFee:      shippingFee,
		CancelAt:         nil,
		PayAt:            nil,
		CreatedAt:        createdAt,
		UpdatedAt:        time.Now(),
		DeletedAt:        nil,
	}
}

func ToCheckoutResponse(
	invoiceNumber string,
	rekeningBank string,
	rekeningNumber string,
	rekeningName string,
	orderDate time.Time,
	totalTransaction uint,
	shippingFee uint,
	totalPay uint,
	address Address,
	products []Product,
) CheckoutResponse {
	return CheckoutResponse{
		InvoiceNumber:    invoiceNumber,
		RekeningBank:     rekeningBank,
		RekeningNumber:   rekeningNumber,
		RekeningName:     rekeningName,
		OrderDate:        orderDate,
		Address:          address,
		Product:          products,
		TotalTransaction: totalTransaction,
		ShippingFee:      shippingFee,
		TotalPay:         totalPay,
	}
}

func ToTransactionProduct(
	data products.Products,
	priceOnCart uint,
	qty uint,

) Product {
	return Product{
		Title:             data.Title,
		Price:             priceOnCart,
		Quantity:          qty,
		TotalProductPrice: priceOnCart * qty,
	}
}

func ToAddressResponse(address address.TransactionAddress) Address {
	return Address{
		Name:        address.Name,
		PhoneNumber: address.PhoneNumber,
		Street:      address.Street,
		City:        address.City,
		Province:    address.Province,
		District:    address.District,
		PostalCode:  address.PostalCode,
		AddressType: address.AddressType,
	}
}
