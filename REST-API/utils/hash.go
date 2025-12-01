package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) { // returns Hashed password

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// byte slice + hash intensity

	return string(bytes), err // convert byte slice to string, and return error
}

// compares input password to hashed password to see if the input could have generated the hash
func CheckPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil // will return false if password is invalid - will yield false if invalid

}
