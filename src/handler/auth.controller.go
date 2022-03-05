package handler

import (
	"net/http"

	"github.com/adiatma85/go-tutorial-gorm/src/helper"
	"github.com/adiatma85/go-tutorial-gorm/src/service"
	"github.com/adiatma85/go-tutorial-gorm/src/validator"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtHelper   helper.JwtHelper
}

func NewAuthHandler(authService service.AuthService, jwtHelper helper.JwtHelper) AuthController {
	return &authController{
		authService,
		jwtHelper,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginRequest validator.LoginValidator

	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		response := helper.BuildFailedResponse("failed request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = c.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response := helper.BuildFailedResponse("failed request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// user service
}

func (c *authController) Register(ctx *gin.Context) {
	var registerRequest validator.RegisterValidator

	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		response := helper.BuildFailedResponse("failed request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// user service create
}
