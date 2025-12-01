package user

import {
	"errors"
	"fmt"
	"time"
}

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time // time is a struct type provided from the time package
}

func (u *user) OutputUserDetails() { // receive function qualifier in front of function.
	//	(*u).firstName  this is the format that would be required, but GO enables the below as a shortcut
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *user) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if FirstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("All fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
