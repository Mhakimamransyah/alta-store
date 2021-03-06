package user

//Service outgoing port for user
type Service interface {
	//FindUserByID If data not found will return nil without error
	FindUserByID(id int) (*User, error)

	//FindUserByEmailAndPassword If data not found will return nil
	FindUserByEmail(email string) (*User, error)

	//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
	FindAllUser(skip int, rowPerPage int) ([]User, error)

	//InsertUser Insert new User into storage
	InsertUser(insertUserSpec InsertUserSpec) error
}

//Repository ingoing port for user
type Repository interface {
	//FindUserByID If data not found will return nil without error
	FindUserByID(id int) (*User, error)

	//FindUserByEmailAndPassword If data not found will return nil
	FindUserByEmail(email string) (*User, error)

	//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
	FindAllUser(skip int, rowPerPage int) ([]User, error)

	//InsertUser Insert new User into storage
	InsertUser(user User) error
}

type Util interface {
	EncryptPassword(string) ([]byte, error)

	ComparePassword(string, string) bool
}
