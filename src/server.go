package src

import (
	"fmt"
	"net/http"

	"github.com/adiatma85/go-tutorial-gorm/src/config"
	"github.com/adiatma85/go-tutorial-gorm/src/handler"
	"github.com/adiatma85/go-tutorial-gorm/src/helper"
	"github.com/adiatma85/go-tutorial-gorm/src/repository"
	"github.com/adiatma85/go-tutorial-gorm/src/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB                  = config.InitializeDatabase()
	userRepo    repository.UserRepository = repository.NewUserRepository(db)
	jwtHelper   helper.JwtHelper          = helper.NewJwtHelper()
	authService service.AuthService       = service.NewAuthService(userRepo)
	authHandler handler.AuthController    = handler.NewAuthHandler(authService, jwtHelper)
)

// Initialize the Server
func Run() {
	// Route
	r := gin.Default()
	v1Route := r.Group("api/v1")
	// Initialize routes
	baseRoutes := v1Route.Group("")
	{
		baseRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "V1 API",
				"data": nil,
				"error": nil
			})
		})
	}
	authRoutes := v1Route.Group("auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("register", authHandler.Register)
	}

	// Running the server
	port := fmt.Sprint(":", viper.GetInt("port"))
	r.Run(port)
}
