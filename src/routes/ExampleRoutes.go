package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExampleRoutes(router *gin.RouterGroup) {
	base := router.Group("example")
	{
		base.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"hit":     "Example",
			})
		})
	}
}