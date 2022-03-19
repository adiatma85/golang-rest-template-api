package v1

import (
	"fmt"

	"github.com/adiatma85/golang-rest-template-api/internal/api/handler"
	"github.com/adiatma85/golang-rest-template-api/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

// V1 Router
func Setup() *gin.Engine {
	app := gin.New()

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

	// Routes for v1
	v1Route := app.Group("/api/v1")

	// AuthGroup with "auth" prefix
	authGroup := v1Route.Group("auth")
	authHandler := handler.GetAuthHandler()
	{
		authGroup.POST("login", authHandler.AuthLogin)
		authGroup.POST("register", authHandler.AuthRegister)
	}

	// UserGroup with "user" prefix
	userGroup := v1Route.Group("users")
	userHandler := handler.GetUserHandler()
	{
		userGroup.GET("", middleware.AuthJWT(), userHandler.GetAllUser)
		userGroup.POST("", middleware.AuthJWT(), userHandler.CreateUser)
		userGroup.GET("query", middleware.AuthJWT(), userHandler.QueryUsers)
		userGroup.GET(":userId", middleware.AuthJWT(), userHandler.GetSpecificUser)
		userGroup.PUT(":userId", middleware.AuthJWT(), userHandler.UpdateSpecificUser)
		userGroup.DELETE(":userId", middleware.AuthJWT(), userHandler.DeleteSpecificUser)
		userGroup.DELETE("multi", middleware.AuthJWT(), userHandler.DeleteUsersWithIds)
	}

	// ProductGroup
	productrGroup := v1Route.Group("products")
	productHandler := handler.GetProductHandler()
	{
		productrGroup.GET("", middleware.AuthJWT(), productHandler.GetAllProduct)
		productrGroup.POST("", middleware.AuthJWT(), productHandler.CreateProduct)
		productrGroup.GET("query", middleware.AuthJWT(), productHandler.QueryProducts)
		productrGroup.GET(":productId", middleware.AuthJWT(), productHandler.GetSpecificProduct)
		productrGroup.PUT(":productId", middleware.AuthJWT(), productHandler.UpdateSpecificProduct)
		productrGroup.DELETE(":productId", middleware.AuthJWT(), productHandler.DeleteSpecificProduct)
		productrGroup.DELETE("multi", middleware.AuthJWT(), productHandler.DeleteProductsWithIds)
	}

	return app
}
