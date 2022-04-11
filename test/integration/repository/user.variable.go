package repository

import (
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
)

// user variable that need to be exist in database before
var users []models.User = []models.User{
	{
		Model:    models.Model{},
		Name:     "Ivan",
		Email:    "ivan@example.com",
		Password: "Password",
		Product:  []models.Product{},
	},
	{
		Name:     "Random User",
		Password: "Random Password",
		Email:    "random_password@example.com",
	},
}

// user that will be inserted in testing
var willBeUser models.User = models.User{
	Model:    models.Model{},
	Name:     "Korenia",
	Email:    "korenia@example.com",
	Password: "Password",
	Product:  []models.Product{},
}
