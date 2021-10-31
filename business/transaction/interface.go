package transaction

type Service interface {
	Checkout(CheckoutSpec) (*CheckoutResponse, error)

	FindTransactionByInvoice(invoiceNumber string, UserID uint) (*CheckoutResponse, error)

	GetAllTransaction(userID uint) ([]*CheckoutResponse, error)
}

type Repository interface {
	CreateTransaction(transaction Transaction) error

	FindTransactionByInvoice(invoiceNumber string) (*Transaction, error)

	FindAllTransaction(listCartID []uint) ([]Transaction, error)

	UpdateTransactionStatus(invoiceNumber, status string) error
}
