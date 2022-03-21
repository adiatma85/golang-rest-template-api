package api

import (
	"fmt"

	v1 "github.com/adiatma85/golang-rest-template-api/internal/api/router/v1"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/config"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/db"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/gin-gonic/gin"
)

// Set configuration
func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)

}

// Initialize repository
func initializeDbRepository() {
	repository.InitializeDb(db.GetDB())
}

// Run the new API with designated configuration
func Run(configPath string) {
	if configPath == "" {
		configPath = "config.yaml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	initializeDbRepository()

	// Routing
	web := v1.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
