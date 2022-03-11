package users

import (
	"time"

	"gorm.io/gorm"
)

// Struct for UserRole Models
type UserRole struct {
	gorm.Model
	UserID   uint64 `gorm:"column:user_id;unique_index:user_role;not null;" json:"user_id"`
	RoleName string `gorm:"column:role_name;not null;" json:"role_name"`
}

// Renew Created_at and Updated_at before creating
func (m *UserRole) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// Renew Created_at and Updated_at before updating
func (m *UserRole) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
