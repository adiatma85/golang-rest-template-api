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
	Cloudinary    StorageCloudinary
	Server        ServerConnection
}

// Struct of Database Configuration instance.
type DatabaseConfiguration struct {
	Driver       string `mapstructure:"DATABASE_DRIVER"`
	Dbname       string `mapstructure:"DATABASE_NAME"`
	Username     string `mapstructure:"DATABASE_USERNAME"`
	Password     string `mapstructure:"DATABASE_PASSWORD"`
	Host         string `mapstructure:"DATABASE_HOST"`
	Port         string `mapstructure:"DATABASE_PORT"`
	MaxLifetime  int    `mapstructure:"DATABASE_MAX_LIFETIME"`
	MaxOpenConns int    `mapstructure:"DATABASE_MAX_OPEN_CONNS"`
	MaxIdleConns int    `mapstructure:"DATABASE_MAX_IDLE_CONNS"`
}

// Struct of Database for Testing Configuration instance
type DatabaseTestConfiguration struct {
	Driver       string `mapstructure:"DATABASE_TEST_DRIVER"`
	Dbname       string `mapstructure:"DATABASE_TEST_NAME"`
	Username     string `mapstructure:"DATABASE_TEST_USERNAME"`
	Password     string `mapstructure:"DATABASE_TEST_PASSWORD"`
	Host         string `mapstructure:"DATABASE_TEST_HOST"`
	Port         string `mapstructure:"DATABASE_TEST_PORT"`
	MaxLifetime  int    `mapstructure:"DATABASE_TEST_MAX_LIFETIME"`
	MaxOpenConns int    `mapstructure:"DATABASE_TEST_MAX_OPEN_CONNS"`
	MaxIdleConns int    `mapstructure:"DATABASE_TEST_MAX_IDLE_CONNS"`
}

// Struct of Cloudinary Storage Configuration instance
type StorageCloudinary struct {
	CloudName    string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	ApiKey       string `mapstructure:"CLOUDINARY_API_KEY"`
	ApiSecret    string `mapstructure:"CLOUDINARY_API_SECRET"`
	UploadFolder string `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
}

// Struct of Server Configuration instance.
type ServerConnection struct {
	Port        string `mapstructure:"SERVER_PORT"`
	Secret      string `mapstructure:"SERVER_SECRET"`
	Mode        string `mapstructure:"SERVER_MODE"`
	Name        string `mapstructure:"SERVER_NAME"`
	ExpiresHour int64  `mapstructure:"SERVER_EXPIRES_HOUR"`
}

// Setup the configuration
func Setup(configPath string) {
	var (
		databaseConfiguration     DatabaseConfiguration
		databaseTestConfiguration DatabaseTestConfiguration
		cloudinaryConfiguration   StorageCloudinary
		serverConfiguration       ServerConnection
	)

	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	unmarshalConfiguration(&databaseConfiguration)
	unmarshalConfiguration(&databaseTestConfiguration)
	unmarshalConfiguration(&cloudinaryConfiguration)
	unmarshalConfiguration(&serverConfiguration)

	configuration := Configuration{
		Database:      databaseConfiguration,
		Database_Test: databaseTestConfiguration,
		Cloudinary:    cloudinaryConfiguration,
		Server:        serverConfiguration,
	}

	Config = &configuration
}

// Helper to unmarshal
func unmarshalConfiguration(configuration interface{}) {
	err := viper.Unmarshal(configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

// GetConfig return the configuration instance
func GetConfig() *Configuration {
	return Config
}
