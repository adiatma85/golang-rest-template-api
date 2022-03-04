package src

import (
	"fmt"

	"github.com/adiatma85/go-tutorial-gorm/src/config"
	"github.com/adiatma85/go-tutorial-gorm/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	// Configuration
	config.Initialize()
	// db := config.InitializeDatabase()

	// Route
	r := gin.Default()
	v1Route := r.Group("api/v1")
	// Initialize routes
	routes.BaseRoutes(v1Route)
	routes.AuthRoutes(v1Route)
	routes.ExampleRoutes(v1Route)

	// Running the server
	port := fmt.Sprint(":", viper.GetInt("port"))
	r.Run(port)
}
