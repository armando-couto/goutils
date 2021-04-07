package goutils

import (
	"golang.org/x/crypto/bcrypt"
)

/*
	Hashing the password with the default cost of 10
*/
func ConvertPassword(password string) string {
	passwordByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

/*
	Comparing the password with the hash
*/
func ToComparePassword(password1, password2 string) error {
	hashedPassword := []byte(password1)
	password := []byte(password2)
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
