package middleware

import (
	"net/http"
	"strings"

	"github.com/adiatma85/go-tutorial-gorm/pkg/crypto"
	"github.com/adiatma85/go-tutorial-gorm/pkg/response"
	"github.com/gin-gonic/gin"
)

// Func to authorizing jwt token
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.Split(authHeader, " ")[1]
		if token == "" {
			response := response.BuildFailedResponse("no token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		jwtHelper := crypto.GetJWTCrypto()
		isValid := jwtHelper.ValidateToken(token)
		if !isValid {
			response := response.BuildFailedResponse("token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
