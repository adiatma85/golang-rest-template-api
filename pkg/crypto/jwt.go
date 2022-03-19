package crypto

import (
	"fmt"
	"time"

	"github.com/adiatma85/golang-rest-template-api/internal/pkg/config"
	"github.com/golang-jwt/jwt"
)

var jwtHelper *jwtCryptoHelper

// Contract fot JWT Crypto Helper
type JWTCryptoHelper interface {
	GenerateToken(UserId string) (string, error)
	ValidateToken(tokenString string) (bool, error)
}

// Struct for jwt custom claim
type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Struct for JWTHelper
type jwtCryptoHelper struct {
}

// Func to initialize new jwt crypto helper
func GetJWTCrypto() JWTCryptoHelper {
	if jwtHelper == nil {
		jwtHelper = &jwtCryptoHelper{}
	}
	return jwtHelper
}

// Func to Generate Token with User ID as main issuer
func (helper *jwtCryptoHelper) GenerateToken(UserID string) (string, error) {
	serverConfiguration := config.GetConfig().Server
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(serverConfiguration.ExpiresHour)).Unix(),
			Issuer:    serverConfiguration.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(serverConfiguration.Secret))
	if err != nil {
		return err.Error(), err
	}
	return t, nil
}

// Func to validate token
func (helper *jwtCryptoHelper) ValidateToken(tokenString string) (bool, error) {
	serverConfiguration := config.GetConfig().Server
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(serverConfiguration.Secret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
