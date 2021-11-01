package auth

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	Login(email string, password string) (string, error)
}

type UtilPassword interface {
	EncryptPassword(string) ([]byte, error)
	ComparePassword(string, string) bool
}
