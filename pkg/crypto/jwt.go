package crypto

import (
	"fmt"
	"time"

	"github.com/adiatma85/go-tutorial-gorm/internal/pkg/config"
	"github.com/golang-jwt/jwt"
)

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Func to Generate Token with User ID as main issuer
func GenerateToken(UserID string) (string, error) {
	serverConfiguration := config.GetConfig().Server
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    serverConfiguration.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(serverConfiguration.Secret))
	if err != nil {
		return "error creating token", err
	}
	return t, nil
}

// Func to validate token
func ValidateToken(tokenString string) bool {
	serverConfiguration := config.GetConfig().Server
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(serverConfiguration.Secret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
