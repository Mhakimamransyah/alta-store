package cart

//Service outgoing port for user
type Service interface {
	//AddToCart create new cart if no cart with status active and create cart detail
	AddToCart(addToCartSpec AddToCartSpec) error
	// //FindUserByID If data not found will return nil without error
	// FindUserByID(id int) (*User, error)

	// //FindUserByEmailAndPassword If data not found will return nil
	// FindUserByEmail(email string) (*User, error)

	// //FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
	// FindAllUser(skip int, rowPerPage int) ([]User, error)

	// //InsertUser Insert new User into storage
	// InsertUser(insertUserSpec InsertUserSpec) error

	//UpdateUser if data not found will return error
	// UpdateUser(id int, name string, modifiedBy string, currentVersion int) error
}

//Repository ingoing port for user
type Repository interface {
	//GetActiveCart If data not found will return nil without error
	GetActiveCart(UserID uint) (*Cart, error)

	//CreateCart create new Cart into storage
	CreateCart(cart Cart) error

	//InsertCartDetail create new cart detail into storage
	InsertCartDetail(cartDetail CartDetail) error

	// //FindUserByEmailAndPassword If data not found will return nil
	// FindUserByEmail(email string) (*User, error)

	// //FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
	// FindAllUser(skip int, rowPerPage int) ([]User, error)

	// //InsertUser Insert new User into storage
	// InsertUser(user User) error

	// //UpdateUser if data not found will return error
	// // UpdateUser(user User, currentVersion int) error
}
