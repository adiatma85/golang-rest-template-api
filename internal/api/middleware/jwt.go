package middleware

import (
	"net/http"
	"strings"

	"github.com/adiatma85/golang-rest-template-api/pkg/crypto"
	"github.com/adiatma85/golang-rest-template-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// Func to authorizing jwt token
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := response.BuildFailedResponse("no token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token := strings.Split(authHeader, " ")[1]
		jwtHelper := crypto.GetJWTCrypto()
		isValid, err := jwtHelper.ValidateToken(token)
		if !isValid {
			response := response.BuildFailedResponse("token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
