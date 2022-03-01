package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseRoutes(router *gin.Engine) {
	base := router.Group("")
	{
		base.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		})
	}
}
