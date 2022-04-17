package api

import (
	"fmt"

	v1 "github.com/adiatma85/golang-rest-template-api/internal/api/router/v1"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/config"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

// Set configuration
// Change this func to "exported"  to make Test package can access it
func SetConfiguration(configPath string) {
	// Setup config from path
	// Default is .env in root folder
	config.Setup(configPath)
	// Calling setup db
	db.SetupDB()
	// Calling cloudinary storage
	// config.InitializeCloudinary()
	gin.SetMode(config.GetConfig().Server.Mode)

}

// Run the new API with designated configuration
func Run(configPath string) {
	if configPath == "" {
		configPath = ".env"
	}
	SetConfiguration(configPath)
	conf := config.GetConfig()

	// Routing
	web := v1.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
