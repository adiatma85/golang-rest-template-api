package api

import (
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/config"
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

// Set configuration
func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

// Run the new API with designated configuration
func Run(configPath string) {
	if configPath == "" {
		configPath = "config.yaml"
	}
	setConfiguration(configPath)
}
