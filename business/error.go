package business

import "errors"

var (
	//ErrInternalServerError Error caused by system error
	ErrInternalServerError = errors.New("internal Server Error")

	//ErrHasBeenModified Error when update item that has been modified
	ErrHasBeenModified = errors.New("data has been modified")

	//ErrNotFound Error when item is not found
	ErrNotFound = errors.New("data was not found")

	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("given spec is not valid")

	//ErrLogin Error when email or password is wrong
	ErrLogin = errors.New("email or password is incorrect")

	//ErrLogin Error when email or password is wrong
	ErrLoginAdmins = errors.New("Username or password is incorrect")

	//ErrRegister Error if duplicate email
	ErrRegister = errors.New("email already registered")

	//ErrAddToCart Error invalid on do parameter
	ErrAddToCart = errors.New("product not found, cannot use subtraction")

	// ErrUnauthorized Error when users try to modify / deleted data that not belongs to him
	ErrUnauthorized = errors.New("unauthorized action")

	//ErrActiveCartNotFound Error if active cart not found
	ErrActiveCartNotFound = errors.New("active cart not found")

	//ErrCartDetailEmpty Error if no product on active cart
	ErrCartDetailEmpty = errors.New("no product on cart to checkout")

	//ErrAddressNotFound Error if address not found
	ErrAddressNotFound = errors.New("address not found")

	//ErrProductNotFound Error if product not found
	ErrProductNotFound = errors.New("product not found")

	//ErrProductOOS error if stock product not enough
	ErrProductOOS = errors.New("insufficient product stock")

	//ErrTransactionNotFound error if transaction not found
	ErrTransactionNotFound = errors.New("transaction not found")

	//ErrTransactionAccess error if try to get transaction another user
	ErrTransactionAccess = errors.New("invalid invoice number")
)
