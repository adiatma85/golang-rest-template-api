package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go"
)

var CloudinaryInstance *cloudinary.Cloudinary

func GetCloudinaryInstance() *cloudinary.Cloudinary {
	return CloudinaryInstance
}

// Initialize Cloudinary
func InitializeCloudinary() {
	config := GetConfig()
	// Create new instance of cloudinary
	cld, err := cloudinary.NewFromParams(config.Cloudinary.CloudName, config.Cloudinary.ApiKey, config.Cloudinary.ApiSecret)
	if err != nil {
		log.Fatalf("Unable to initialize cloudinary storage, %v", err)
	}

	CloudinaryInstance = cld
}
