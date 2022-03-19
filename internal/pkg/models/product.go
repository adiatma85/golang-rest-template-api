package models

// Struct that define Product models
type Product struct {
	Model
	Name   string `gorm:"not null;type:varchar(100)" json:"name" validation:"name"`
	Price  uint64 `gorm:"type:bigint" json:"price" validation:"price"`
	UserId int64  `gorm:"not null" json:"-" validation:"user_id"`
	User   User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
