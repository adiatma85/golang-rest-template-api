package db

import (
	"fmt"
	"time"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/config"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

// Database instance
type Database struct {
	*gorm.DB
}

// SetupDB is a function to open connection to database
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	// Viper Config
	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	// Gorm config
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local"), gormConfig)
		if err != nil {
			fmt.Println("db err:", err)
		}
	case "postgres":
		db, err = gorm.Open(postgres.Open("host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password), gormConfig)
		if err != nil {
			fmt.Println("db err:", err)
		}
	}
	// Set up the connection pools
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(configuration.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(configuration.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime))

	DB = db
	migration()
}

// Setup for testing database
func SetupTestingDb(host, username, password, port, database string) {
	// For the sake of simplicity, right now database testing is in mysql
	db, err := gorm.Open(mysql.Open(username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err for testing :", err)
		panic(err.Error())
	}

	DB = db
	migration()
}

// AutoMigrate project models
func migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}

func GetDB() *gorm.DB {
	return DB
}
