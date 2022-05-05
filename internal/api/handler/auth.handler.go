package handler

import (
	"fmt"
	"net/http"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/models"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/repository"
	"github.com/adiatma85/golang-rest-template-api/internal/pkg/validator"
	"github.com/adiatma85/golang-rest-template-api/pkg/crypto"
	"github.com/adiatma85/golang-rest-template-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

// Local variable
var (
	authHandler *AuthHandler
)

// Struct to implement contract of authHandler
type AuthHandler struct{}

// Contract of Auth Handler
type AuthHandlerInterface interface {
	AuthLogin(c *gin.Context)
	AuthRegister(c *gin.Context)
}

// Func to return Auth Handler instance
func GetAuthHandler() AuthHandlerInterface {
	if authHandler == nil {
		authHandler = &AuthHandler{}
	}
	return authHandler
}

// Func to handle Auth Login
func (handler *AuthHandler) AuthLogin(c *gin.Context) {
	var loginRequest validator.LoginRequest
	err := c.ShouldBind(&loginRequest)

	// Error when binding in validator
	if err != nil {
		response := response.BuildFailedResponse("failed to login", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userRepo := repository.GetUserRepository()
	// If user doesn't exist
	if user, err := userRepo.GetByEmail(loginRequest.Email); err != nil {
		response := response.BuildFailedResponse("failed to login", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		// Wrong password
		passwordHelper := crypto.GetPasswordCryptoHelper()
		if !passwordHelper.ComparePassword(user.Password, []byte(loginRequest.Password)) {
			response := response.BuildFailedResponse("wrong credential", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		// Correct password
		tokenHelper := crypto.GetJWTCrypto()
		token, err := tokenHelper.GenerateToken(fmt.Sprint(user.ID))
		if err != nil {
			response := response.BuildFailedResponse("wrong credential", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		response := response.BuildSuccessResponse("success login", map[string]interface{}{
			"token": token,
		})
		c.JSON(http.StatusOK, response)
		return
	}
}

// Func to handle Auth Register
func (handler *AuthHandler) AuthRegister(c *gin.Context) {
	// var uploadSecureUrl = ""
	var registerRequest validator.RegisterRequest
	err := c.ShouldBind(&registerRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to register", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// avatarImageFile, isExist := helpers.GinFileHandlerFunc(c, "avatar")

	// Jika memang ada, maka upload image tersebut melalui cloudinary
	// if isExist {
	// 	uploadSecureUrl, err = helpers.CloudinaryImageUploadHelper(avatarImageFile)
	// 	if err != nil {
	// 		response := response.BuildFailedResponse("failed to register new user", err.Error())
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// }

	userRepo := repository.GetUserRepository()
	passwordHelper := crypto.GetPasswordCryptoHelper()
	userModel := &models.User{}

	// smapping the struct
	smapping.FillStruct(userModel, smapping.MapFields(&registerRequest))
	userModel.Password, _ = passwordHelper.HashAndSalt([]byte(registerRequest.Password))
	// userModel.Avatar = uploadSecureUrl

	if newUser, err := userRepo.Create(*userModel); err != nil {
		response := response.BuildFailedResponse("failed to register", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		tokenHelper := crypto.GetJWTCrypto()
		token, err := tokenHelper.GenerateToken(fmt.Sprint(newUser.ID))
		if err != nil {
			response := response.BuildFailedResponse("failed to generate token", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := response.BuildSuccessResponse("success register new user", map[string]interface{}{
			"token": token,
		})
		c.JSON(http.StatusOK, response)
		return
	}
}
