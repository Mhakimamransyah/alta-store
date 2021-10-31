package transaction

type Service interface {
	Checkout(CheckoutSpec) (*CheckoutResponse, error)
}

type Repository interface {
	CreateTransaction(transaction Transaction) error
}
