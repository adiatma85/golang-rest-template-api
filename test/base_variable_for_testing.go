package test

import (
	// "github.com/adiatma85/golang-rest-template-api/internal/api"

	"fmt"

	"github.com/adiatma85/golang-rest-template-api/internal/api"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/config"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
)

// Helper for database params
var (

	// Get the configuration in here

	// Database configuration
	// Change configuration in here
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

// Initialize func to call configuration
func SetupInitialize(path string) {
	api.SetConfiguration(path)
	configuration := config.GetConfig()
	fmt.Println(configuration.DatabaseTest.Dbname)
	Database = configuration.DatabaseTest.Dbname
	Username = configuration.DatabaseTest.Username
	Password = configuration.DatabaseTest.Password
	Host = configuration.DatabaseTest.Host
	Port = configuration.DatabaseTest.Port
}

// TeardownHelper
func TearDownHelper() {
	for _, model := range Models {
		db.GetDB().Migrator().DropTable(model)
	}
}
