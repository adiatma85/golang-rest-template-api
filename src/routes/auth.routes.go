package routes

import (
	"github.com/adiatma85/go-tutorial-gorm/src/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {

	// Authcontroller
	authController := controller.NewAuthController()

	// AuthRoutes
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Login)
	}
}
