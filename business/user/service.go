package user

import (
	"altaStore/business"
	"altaStore/util/validator"
	"fmt"
	"time"
)

//InsertUserSpec create user spec
type InsertUserSpec struct {
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	PhoneNumber string `validate:"required,number"`
	Password    string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository  Repository
	utilService Util
}

//NewService Construct user service object
func NewService(repository Repository, util Util) Service {
	return &service{
		repository,
		util,
	}
}

//FindUserByID Get user by given ID, return nil if not exist
func (s *service) FindUserByID(id int) (*User, error) {
	x, err := s.repository.FindUserByID(id)
	fmt.Println("TTTTTTTTTTTTTTTTTTTTTTTTTT")
	fmt.Println(err)
	fmt.Println("__________________________")
	fmt.Println(x)
	fmt.Println("TTTTTTTTTTTTTTTTTTTTTTTTTT")
	// return s.repository.FindUserByID(id)
	return x, err
}

//FindUserByEmailAndPassword Get user by given ID, return nil if not exist
func (s *service) FindUserByEmail(email string) (*User, error) {
	return s.repository.FindUserByEmail(email)
}

//FindAllUser Get all users , will be return empty array if no data or error occured
func (s *service) FindAllUser(skip int, rowPerPage int) ([]User, error) {

	user, err := s.repository.FindAllUser(skip, rowPerPage)
	fmt.Println("TTTTTTTTTTTTTTTTTTTTTTTTTT")
	fmt.Println(err)
	fmt.Println("__________________________")
	fmt.Println(user)
	fmt.Println("TTTTTTTTTTTTTTTTTTTTTTTTTT")
	if err != nil {
		return []User{}, err
	}

	return user, err
}

//InsertUser Create new user and store into database
func (s *service) InsertUser(insertUserSpec InsertUserSpec) error {
	err := validator.GetValidator().Struct(insertUserSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}
	hashPassword, errorEncrypt := s.utilService.EncryptPassword(insertUserSpec.Password)
	if errorEncrypt != nil {
		return errorEncrypt
	}
	user := NewUser(
		insertUserSpec.Name,
		insertUserSpec.Email,
		insertUserSpec.PhoneNumber,
		string(hashPassword),
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return business.ErrRegister
	}

	return nil
}
