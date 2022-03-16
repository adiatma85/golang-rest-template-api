package v1

import (
	"fmt"

	"github.com/adiatma85/go-tutorial-gorm/internal/api/handler"
	"github.com/adiatma85/go-tutorial-gorm/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

// V1 Router
func Setup() *gin.Engine {
	app := gin.New()

	// Middlewares
	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	// NEED TO ADD JWT MIDDLEWARE

	// Routes for v1
	v1Route := app.Group("/api/v1")

	// AuthGroup with "auth" prefix
	authGroup := v1Route.Group("auth")
	{
		authGroup.POST("login", handler.AuthLoginHandler)
		authGroup.POST("register", handler.AuthRegisterHandler)
	}

	// UserGroup with "user" prefix
	userGroup := v1Route.Group("users")
	{
		userGroup.GET("", handler.GetAllUser)
		userGroup.POST("", handler.CreateUser)
		userGroup.GET("query", handler.QueryUsers)
		userGroup.GET(":userId", handler.GetSpecificUser)
		userGroup.PUT(":userId", handler.UpdateSpecificUser)
		userGroup.DELETE(":userId", handler.DeleteSpecificUser)
		userGroup.DELETE("multi", handler.DeleteUsersWithIds)
	}

	// ProductGroup
	productrGroup := v1Route.Group("products")
	{
		productrGroup.GET("", handler.GetAllProduct)
		productrGroup.POST("", handler.CreateProduct)
		productrGroup.GET("query", handler.QueryProducts)
		productrGroup.GET(":productId", handler.GetSpecificProduct)
		productrGroup.PUT(":productId", handler.UpdateSpecificProduct)
		productrGroup.DELETE(":productId", handler.DeleteSpecificProduct)
		productrGroup.DELETE("multi", handler.DeleteProductsWithIds)
	}

	return app
}