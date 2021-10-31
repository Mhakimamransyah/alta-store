package business

import "errors"

var (
	//ErrInternalServerError Error caused by system error
	ErrInternalServerError = errors.New("Internal Server Error")

	//ErrHasBeenModified Error when update item that has been modified
	ErrHasBeenModified = errors.New("Data has been modified")

	//ErrNotFound Error when item is not found
	ErrNotFound = errors.New("Data was not found")

	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("Given spec is not valid")

	//ErrLogin Error when email or password is wrong
	ErrLogin = errors.New("Email or password is incorrect")

	//ErrLogin Error when email or password is wrong
	ErrLoginAdmins = errors.New("Username or password is incorrect")

	//ErrRegister Error if duplicate email
	ErrRegister = errors.New("Email already registered")

	//ErrAddToCart Error invalid on do parameter
	ErrAddToCart = errors.New("Product not found, cannot use subtraction")

	// ErrUnauthorized Error when users try to modify / deleted data that not belongs to him
	ErrUnauthorized = errors.New("Unauthorized action")
)
