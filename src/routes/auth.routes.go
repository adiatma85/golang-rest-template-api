package routes

import (
	"github.com/adiatma85/go-tutorial-gorm/src/handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {

	// Authcontroller
	authController := handler.NewAuthHandler()

	// AuthRoutes
	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Login)
	}
}
