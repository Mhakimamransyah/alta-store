package response

import (
	"altaStore/business/transaction"
	"time"
)

type allTransactionResponse struct {
	Trx []transactionResponse
}

type transactionResponse struct {
	InvoiceNumber    string
	RekeningBank     string
	RekeningNumber   string
	RekeningName     string
	OrderDate        time.Time
	Address          transaction.Address
	Product          []transaction.Product
	TotalTransaction uint
	ShippingFee      uint
	TotalPay         uint
}

func NewTransactionResponse(transaction transaction.CheckoutResponse) transactionResponse {
	var trx transactionResponse

	trx.InvoiceNumber = transaction.InvoiceNumber
	trx.RekeningBank = transaction.RekeningBank
	trx.RekeningNumber = transaction.RekeningNumber
	trx.RekeningName = transaction.RekeningName
	trx.OrderDate = transaction.OrderDate
	trx.Address = transaction.Address
	trx.Product = transaction.Product
	trx.TotalTransaction = transaction.TotalTransaction
	trx.ShippingFee = transaction.ShippingFee
	trx.TotalPay = transaction.TotalPay

	return trx
}
