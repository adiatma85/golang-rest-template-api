package crypto

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

var bcryptPasswordHelper *passwordCryptoHelper

// // Contract fot Password Crypto Helper
type PasswordCryptoHelper interface {
	HashAndSalt(pwd []byte) (string, error)
	ComparePassword(hashPassword string, plainPassword []byte) bool
}

// // Struct to implement Password Crypto helper
type passwordCryptoHelper struct {
}

// // Func to initialize Password Crypto Helper
func GetPasswordCryptoHelper() PasswordCryptoHelper {
	if bcryptPasswordHelper == nil {
		bcryptPasswordHelper = &passwordCryptoHelper{}
	}
	return bcryptPasswordHelper
}

// Generate Hash from byte Password
func (helper *passwordCryptoHelper) HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

// Compare Password between Hashed ones and Plain
// If Password match return true
// If Password doesn't match return false
func (helper *passwordCryptoHelper) ComparePassword(hashPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}
