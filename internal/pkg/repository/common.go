package repository

import (
	"errors"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/db"
	"gorm.io/gorm"
)

// Common function to create in db
func Create(value interface{}) error {
	return db.GetDB().Create(value).Error
}

// Common function to save in db
func Save(value interface{}) error {
	return db.GetDB().Save(value).Error
}

// Common function to get the first row
// Associations mean its relation to other
func First(where interface{}, out interface{}, associations []string) (notFound bool, err error) {
	db := db.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	err = db.Where(where).First(out).Error
	if err != nil {
		notFound = errors.Is(err, gorm.ErrRecordNotFound)
	}
	return
}

// Common function to update in db
func Update(where, value interface{}) error {
	return db.GetDB().Model(where).Updates(value).Error
}

// Common function to find in db
func Find(where interface{}, out interface{}, associations []string, orders ...string) error {
	db := db.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Common function to scan in db
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = db.GetDB().Model(model).Where(where).Scan(out).Error
	if err != nil {
		notFound = errors.Is(err, gorm.ErrRecordNotFound)
	}
	return
}

// Common function to scanlist in db
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	db := db.GetDB().Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// Common function to delete by model in db
func DeleteByModel(model interface{}) (count int64, err error) {
	db := db.GetDB().Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Common function to delete by where in db
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := db.GetDB().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Common function to delete by id in db
func DeleteByID(model interface{}, id uint64) (count int64, err error) {
	db := db.GetDB().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Common function to delete by ids (multiple) in db
func DeleteByIDS(model interface{}, ids []uint64) (count int64, err error) {
	db := db.GetDB().Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}
