package handler

import (
	"net/http"

	"github.com/adiatma85/go-tutorial-gorm/pkg/response"
	"github.com/gin-gonic/gin"
)

// Func to handle Auth Login
func AuthLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, response.BuildSuccessResponse("success login", nil))
}

func AuthRegisterHandler(c *gin.Context) {
	c.JSON(http.StatusOK, response.BuildSuccessResponse("success register", nil))
}
