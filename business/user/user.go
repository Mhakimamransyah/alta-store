package user

import "time"

type User struct {
	ID          uint
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

//NewUser create new User
func NewUser(
	name string,
	email string,
	phoneNumber string,
	password string,
	createdAt time.Time) User {

	return User{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
		CreatedAt:   createdAt,
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
}
