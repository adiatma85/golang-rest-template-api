package crypto

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Contract fot Password Crypto Helper
type PasswordCryptoHelper interface {
	HashAndSalt(pwd []byte) (string, error)
	ComparePassword(hashPassword string, plainPassword []byte) bool
}

// Struct to implement Password Crypto helper
type passwordCryptoHelper struct {
}

// Func to initialize Password Crypto Helper
func NewPasswordCryptoHelper() PasswordCryptoHelper {
	return &passwordCryptoHelper{}
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
func (helper *passwordCryptoHelper) ComparePassword(hashPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	return err == nil
}
