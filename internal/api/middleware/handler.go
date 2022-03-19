package middleware

import (
	"net/http"

	"github.com/adiatma85/golang-rest-template-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// No Method Handler global middleware
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, response.BuildFailedResponse("method not permitted", nil))
	}
}

// No Route Handler global middleware
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.BuildFailedResponse("the processing function of the request route was not found", nil))
	}
}
