package test

import (
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
)

// Helper for database params
var (

	// Database configuration
	Database = "go-unit-integration"
	Username = "root"
	Password = "root"
	Host     = "localhost"
	Port     = "3306"

	// Models that involved
	Models = []interface{}{
		&models.User{},
		&models.Product{},
	}
)
