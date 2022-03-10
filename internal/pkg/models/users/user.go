package users

import (
	"time"

	"gorm.io/gorm"
)

// Struct for User Models
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" json:"-"`
	Email    string `gorm:"type:varchar(100);unique;" json:"-"`
	Password string `gorm:"type:varchar(100)" json:"-"`
	// Role     UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Renew Created_at and Updated_at before creating
func (m *User) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// Renew Created_at and Updated_at before updating
func (m *User) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
