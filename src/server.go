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

	// Route
	r := gin.Default()
	// Initialize routes
	routes.BaseRoutes(r)
	routes.ExampleRoutes(r)

	// Running the server
	port := fmt.Sprint(":", viper.GetInt("port"))
	r.Run(port)
}
