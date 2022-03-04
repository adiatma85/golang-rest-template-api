package controller

import (
	"net/http"

	"github.com/adiatma85/go-tutorial-gorm/src/helper"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	respond := helper.BuildSuccessResponse("nice", "login")
	ctx.JSON(http.StatusOK, respond)
}

func (c *authController) Register(ctx *gin.Context) {
	respond := helper.BuildSuccessResponse("nice", "register")
	ctx.JSON(http.StatusOK, respond)
}
