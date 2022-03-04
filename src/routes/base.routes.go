package routes

import (
	"github.com/adiatma85/go-tutorial-gorm/src/controller"
	"github.com/gin-gonic/gin"
)

func BaseRoutes(router *gin.RouterGroup) {

	// baseController
	baseController := controller.NewBaseController()

	baseRoutes := router.Group("")
	{
		baseRoutes.GET("/", baseController.Base)
	}
}
