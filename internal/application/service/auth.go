package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	dayInHours = 24
)

type TokenService interface {
	GenToken(userId string) (string, error)
	VerifyToken(tokenString string) (*Token, error)
}

type Token struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewAuthService() TokenService {
	return &Token{}
}

func (t *Token) GenToken(userId string) (string, error) {
	claims := Token{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * dayInHours)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (t *Token) VerifyToken(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
