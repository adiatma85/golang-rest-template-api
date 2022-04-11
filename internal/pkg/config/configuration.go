package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

// Struct of Configuration instance.
// It include Database and Server configuration
type Configuration struct {
	Database      DatabaseConfiguration
	Database_Test DatabaseTestConfiguration
	Server        ServerConnection
}

// Struct of Database Configuration instance.
type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// Struct of Database for Testing Configuration instance
type DatabaseTestConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// Struct of Server Configuration instance.
type ServerConnection struct {
	Port        string
	Secret      string
	Mode        string
	Name        string
	ExpiresHour int64
}

// Setup the configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	Config = configuration
}

// GetConfig return the configuration instance
func GetConfig() *Configuration {
	return Config
}
