package models

import (
	"github.com/adiatma85/golang-rest-template-api/pkg/crypto"
	"gorm.io/gorm"
)

// User Type
type userType string

const (
	ADMIN userType = "ADMIN"
	USER  userType = "USER"
)

// Struct for User Models
type User struct {
	Model
	Name     string    `gorm:"type:varchar(100)" json:"name" validation:"name"`
	Email    string    `gorm:"type:varchar(100);unique;" json:"email" validation:"email"`
	Password string    `gorm:"type:varchar(100)" json:"-" validation:"password"`
	Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Renew Created_at and Updated_at before creating
func (m *User) BeforeCreate(db *gorm.DB) error {
	passwordHelper := crypto.GetPasswordCryptoHelper()
	hashedPassword, err := passwordHelper.HashAndSalt([]byte(m.Password))
	m.Password = hashedPassword
	return err
}

// Renew Created_at and Updated_at before updating
func (m *User) BeforeUpdate(db *gorm.DB) error {
	if m.Password != "" {
		passwordHelper := crypto.GetPasswordCryptoHelper()
		newHashedPassword, err := passwordHelper.HashAndSalt([]byte(m.Password))
		m.Password = newHashedPassword
		return err
	}
	return nil
}
