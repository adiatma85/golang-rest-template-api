package config

import (
	"fmt"

	"github.com/adiatma85/go-tutorial-gorm/src/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Func to Initialize database
func InitializeDatabase() *gorm.DB {
	databaseHost := viper.GetString("Database.Host")
	databaseUsername := viper.GetString("Database.Username")
	databasePassword := viper.GetString("Database.Password")
	databaseDb := viper.GetString("Database.Database")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUsername, databasePassword, databaseHost, databaseDb)

	// Connecting do database
	fmt.Println(connectionString)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	db.AutoMigrate(&models.User{}, &models.Book{})
	return db
}

// func mysqlSqlString(host, username, passwrod, db string) string {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, passwrod, host, db)
// 	return dsn
// }
