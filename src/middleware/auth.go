package middleware

import (
	"log"
	"net/http"

	"github.com/adiatma85/go-tutorial-gorm/src/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJwt(jwtHelper helper.JwtHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildFailedResponse("failed to process request", "no token provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}

		token := jwtHelper.ValidateToken(authHeader, c)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			response := helper.BuildFailedResponse("error", "token is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
