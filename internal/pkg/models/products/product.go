package products

import (
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/models"
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/models/users"
)

type Product struct {
	models.Model
	Name   string     `gorm:"not null;type:varchar(100)" json:"-"`
	Price  uint64     `gorm:"type:bigint" json:"-"`
	UserID int64      `gorm:"not null" json:"-"`
	User   users.User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
