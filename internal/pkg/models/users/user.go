package users

import (
	"time"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/models"
)

// Struct for User Models
type User struct {
	models.Model
	Username  string   `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	Firstname string   `gorm:"column:firstname;not null;" json:"firstname" form:"firstname"`
	Lastname  string   `gorm:"column:lastname;not null;" json:"lastname" form:"lastname"`
	Password  string   `gorm:"column:hash;not null;" json:"hash"`
	Role      UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
