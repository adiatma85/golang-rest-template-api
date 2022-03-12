package api

import (
	"fmt"

	v1 "github.com/adiatma85/go-tutorial-gorm/internal/api/router/v1"
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
	conf := config.GetConfig()
	web := v1.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
