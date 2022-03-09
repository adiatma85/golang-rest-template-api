package handler

import (
	"net/http"
	"strconv"

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
	userService service.UserService
}

func NewAuthHandler(authService service.AuthService, jwtHelper helper.JwtHelper, userService service.UserService) AuthController {
	return &authController{
		authService,
		jwtHelper,
		userService,
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

	user, _ := c.userService.FindUserByEmail(loginRequest.Email)
	token := c.jwtHelper.GenerateToken(strconv.FormatInt(int64(user.ID), 10))
	user.Token = token
	response := helper.BuildSuccessResponse("succes login", user)

	ctx.JSON(http.StatusOK, response)
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
	user, err := c.userService.CreateUser(registerRequest)
	if err != nil {
		response := helper.BuildFailedResponse("register failed", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	token := c.jwtHelper.GenerateToken(strconv.FormatInt(int64(user.ID), 10))
	user.Token = token
	response := helper.BuildSuccessResponse("succes register", user)

	ctx.JSON(http.StatusOK, response)
}
