package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string //exported
	lastName  string // not exported
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name, and birthdate are required.")
	}
	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "-----",
			createdAt: time.Now(),
		},
	}

}
