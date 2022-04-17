package repository

import "github.com/adiatma85/golang-rest-template-api/internal/pkg/models"

// product variable that need to be exist in database before
var products []models.Product = []models.Product{
	{
		Model:  models.Model{},
		Name:   "Item 1",
		Price:  10000,
		UserId: 1,
	},
	{
		Model:  models.Model{},
		Name:   "Item 2",
		Price:  20000,
		UserId: 1,
	},
	{
		Model:  models.Model{},
		Name:   "Item 3",
		Price:  40000,
		UserId: 2,
	},
}

// product that will be inserted in testing
var willBeProduct = models.Product{
	Model:  models.Model{},
	Name:   "Item Random milik Korenia",
	Price:  5000,
	UserId: 1,
}
