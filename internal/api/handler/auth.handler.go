package handler

import (
	"fmt"
	"net/http"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/repository"
	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/validator"
	"github.com/adiatma85/go-tutorial-gorm/pkg/crypto"
	"github.com/adiatma85/go-tutorial-gorm/pkg/response"
	"github.com/gin-gonic/gin"
)

// Func to handle Auth Login
func AuthLoginHandler(c *gin.Context) {
	var loginRequest validator.LoginRequest
	err := c.ShouldBind(&loginRequest)

	// Error when binding in validator
	if err != nil {
		response := response.BuildFailedResponse("failed to login", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userRepo := repository.GetUserRepository()
	// If user doesn't exist
	if user, err := userRepo.GetByEmail(loginRequest.Email); err != nil {
		response := response.BuildFailedResponse("failed to login", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	} else {
		// Wrong password
		passwordHelper := crypto.GetPasswordCryptoHelper()
		if !passwordHelper.ComparePassword(user.Password, []byte(loginRequest.Password)) {
			response := response.BuildFailedResponse("wrong credential", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Correct password
		tokenHelper := crypto.GetJWTCrypto()
		token, err := tokenHelper.GenerateToken(fmt.Sprint(user.ID))
		if err != nil {
			response := response.BuildFailedResponse("wrong credential", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
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
func AuthRegisterHandler(c *gin.Context) {
	var registerRequest validator.RegisterRequest
	err := c.ShouldBind(&registerRequest)

	if err != nil {
		response := response.BuildFailedResponse("failed to register", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userRepo := repository.GetUserRepository()

	if newUser, err := userRepo.Create(registerRequest); err != nil {
		response := response.BuildFailedResponse("failed to register", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	} else {
		tokenHelper := crypto.GetJWTCrypto()
		token, err := tokenHelper.GenerateToken(fmt.Sprint(newUser.ID))
		if err != nil {
			response := response.BuildFailedResponse("wrong credential", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := response.BuildSuccessResponse("success login", map[string]interface{}{
			"token": token,
		})
		c.JSON(http.StatusOK, response)
		return
	}
}
