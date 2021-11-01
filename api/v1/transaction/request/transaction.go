package request

import "altaStore/business/transaction"

type CheckoutRequest struct {
	UserID      uint `validate:"gt=0"`
	CartID      uint `validate:"gt=0"`
	ShippingFee uint `validate:"gte=0"`
	AddressID   uint `validate:"gte=0"`
}

type Invoice struct {
	invoiceNumber string `validate:"required"`
}

func (req *CheckoutRequest) ToCheckoutSpec() *transaction.CheckoutSpec {

	var checkoutSpec transaction.CheckoutSpec
	checkoutSpec.UserID = req.UserID
	checkoutSpec.CartID = req.CartID
	checkoutSpec.ShippingFee = req.ShippingFee
	checkoutSpec.AddressID = req.AddressID
	return &checkoutSpec
}

func (req *Invoice) ToInvoice() *Invoice {
	var invoice Invoice
	invoice.invoiceNumber = req.invoiceNumber
	return &invoice
}
