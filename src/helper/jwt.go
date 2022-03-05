package helper

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtHelper interface {
	GenerateToken(userId string) string
	ValidateToken(token string, ctx *gin.Context) *jwt.Token
}

type jwtClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtHelper struct {
	secretKey string
	issuer    string
}

// Function to initialize instance of jwt helper
func NewJwtHelper() JwtHelper {
	return &jwtHelper{
		issuer:    "",
		secretKey: "",
	}
}

// Func to generate token from user id
func (s *jwtHelper) GenerateToken(userId string) string {
	claims := &jwtClaim{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// Func to validate token from string. This is used in gin framework
func (s *jwtHelper) ValidateToken(token string, ctx *gin.Context) *jwt.Token {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil
	}
	return t
}
