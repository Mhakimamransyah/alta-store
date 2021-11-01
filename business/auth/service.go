package auth

import (
	business "altaStore/business"
	"altaStore/business/user"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//=============== The implementation of those interface put below =======================
type service struct {
	userService user.Service
	utilService UtilPassword
}

//NewService Construct user service object
func NewService(userService user.Service, utilService UtilPassword) Service {
	return &service{
		userService,
		utilService,
	}
}

//Login by given user Email and Password, return error if not exist
func (s *service) Login(email string, password string) (string, error) {
	user, err := s.userService.FindUserByEmail(email)
	if err != nil {
		return "", business.ErrLogin
	}

	comparePassword := s.utilService.ComparePassword(user.Password, password)

	if !comparePassword {
		return "", business.ErrLogin
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["name"] = user.Name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
