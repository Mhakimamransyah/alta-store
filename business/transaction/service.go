package transaction

//AddToCartSpec create cart and cart detail spec
type CreateTransactionSpec struct {
	CartID      uint `validate:"number"`
	shippingFee uint `validate:"number"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CreateTransaction(CreateTransactionSpec) error {
	return nil
}
