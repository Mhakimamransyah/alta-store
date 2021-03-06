package address

type Service interface {
	InsertAddress(insertAddressSpec InsertAddressSpec) error

	GetAllAddress(userID uint) ([]Address, error)

	GetDefaultAddress(userID uint) (*Address, error)
}

//Repository ingoing port for user
type Repository interface {
	InsertAddress(address Address) error

	GetAllAddress(userID uint) ([]Address, error)

	GetDefaultAddress(userID uint) (*Address, error)

	UpdateDefaultAddress(address Address) error

	GetAddressForTransaction(userID, addressID uint) (*TransactionAddress, error)

	GetAddressByID(addressID uint) (*TransactionAddress, error)
}
