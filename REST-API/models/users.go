package models

// bcrypt package can help with hashing
// go get -u golang.org/x/crypto

import (
	"errors"

	"example.com/REST-API/db"
	"example.com/REST-API/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	u.ID = userID

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email) // will get 1 row or no results given email is unique

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword) // originally updated a copy

	if err != nil {
		return errors.New("Credentials Invalid")
	}

	// if you made it this far, you know a row exists and you have a password
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials Invalid")
	}

	// know credentials are valid so DO NOT return an error
	return nil
}
