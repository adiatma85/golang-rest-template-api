package routes

import (
	"github.com/adiatma85/go-tutorial-gorm/src/handler"
	"github.com/gin-gonic/gin"
)

func BaseRoutes(router *gin.RouterGroup) {

	// baseController
	baseController := handler.NewBaseHandler()

	baseRoutes := router.Group("")
	{
		baseRoutes.GET("/", baseController.Base)
	}
}
