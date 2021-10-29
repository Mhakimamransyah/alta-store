package transaction

import "time"

type Transaction struct {
	ID               uint
	CartID           uint
	InvoiceNumber    string
	Status           string
	TotalTransaction uint
	ShippingFee      uint
	CancelAt         *time.Time
	PayAt            *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
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
