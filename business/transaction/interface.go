package transaction

type Service interface {
	CreateTransaction(CreateTransactionSpec) error
}

type Repository interface {
	CreateTransaction(transaction Transaction) error
}
