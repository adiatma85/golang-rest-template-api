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
	Database string
	Username string
	Password string
	Host     string
	Port     string

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
	fmt.Println(configuration.Database_Test.Dbname)
	Database = configuration.Database_Test.Dbname
	Username = configuration.Database_Test.Username
	Password = configuration.Database_Test.Password
	Host = configuration.Database_Test.Host
	Port = configuration.Database_Test.Port
}

// TeardownHelper
func TearDownHelper() {
	for _, model := range Models {
		db.GetDB().Migrator().DropTable(model)
	}
}
