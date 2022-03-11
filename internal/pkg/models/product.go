package models

// Struct that define Product models
type Product struct {
	Model
	Name   string `gorm:"not null;type:varchar(100)" json:"-"`
	Price  uint64 `gorm:"type:bigint" json:"-"`
	UserID int64  `gorm:"not null" json:"-"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
