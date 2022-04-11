package models

import (
	"time"

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
	UserType userType  `gorm:"type:varchar(10);default:ADMIN" json:"user_type"`
}

// Renew Created_at and Updated_at before creating
func (m *User) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// Renew Created_at and Updated_at before updating
func (m *User) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
